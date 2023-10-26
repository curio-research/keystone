package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/curio-research/keystone/core"
	"github.com/curio-research/keystone/server"
	"github.com/golang-collections/collections/stack"
	"gorm.io/gorm"
)

type MySQLSaveStateHandler struct {
	dbConnection       *gorm.DB
	gameID             string
	schemasToAccessors map[interface{}]*core.TableBaseAccessor[any]
}

// initialize connection mySQL
func newSQLSaveStateHandler(dialector gorm.Dialector, gameID string, schemasToAccessors map[interface{}]*core.TableBaseAccessor[any]) (*MySQLSaveStateHandler, error) {
	db, err := gorm.Open(dialector, gormOpts(gameID))
	if err != nil {
		return nil, err
	}

	handler := &MySQLSaveStateHandler{
		dbConnection:       db,
		gameID:             gameID,
		schemasToAccessors: schemasToAccessors,
	}

	err = handler.initializeDBTables()
	if err != nil {
		return nil, err
	}

	return handler, nil
}

// initialize mySQL tables for saving state updates
func (handler *MySQLSaveStateHandler) initializeDBTables() error {
	db := handler.dbConnection
	if db == nil {
		return fmt.Errorf("db connection is nil")
	}

	// all tables that need to be created
	allSchemas := []any{}
	for schema, _ := range handler.schemasToAccessors {
		if reflect.TypeOf(schema).Kind() != reflect.Pointer {
			return fmt.Errorf("schema %v is not a pointer to the struct", schema)
		}
		allSchemas = append(allSchemas, schema)
	}

	err := db.AutoMigrate(allSchemas...)
	if err != nil {
		return err
	}

	fmt.Println("-> All tables have been created")
	return nil
}

// save state updates to mySQL database
func (handler *MySQLSaveStateHandler) SaveState(tableUpdates []core.TableUpdate) error {
	// process table updates
	tableUpdateOperationsByTable, tableRemovalOperationsByTable := processUpdatesForUpload(tableUpdates)

	// update operations
	for table, updates := range tableUpdateOperationsByTable {
		arr := handler.castToSchemaArray(table, updates)
		if arr != nil {
			tx := handler.dbConnection.Save(arr)
			if tx.Error != nil {
				return tx.Error
			}
		}
	}

	// removal operations
	for table, removals := range tableRemovalOperationsByTable {
		arr := handler.castToSchemaArray(table, removals)
		tx := handler.dbConnection.Delete(arr)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

// given a schema type, use the mapping from tables to cast to an array of that type
func (handler *MySQLSaveStateHandler) castToSchemaArray(schemaType string, vals []interface{}) interface{} {
	var accessor *core.TableBaseAccessor[any]
	for _, schemaAccessor := range handler.schemasToAccessors {
		if strings.Contains(schemaAccessor.Name(), schemaType) {
			accessor = schemaAccessor
			break
		}
	}
	if accessor == nil {
		return nil
	}

	schema := accessor.Type()

	// Use reflection to cast val to the appropriate schema type.
	arrayType := reflect.SliceOf(schema)
	castedValue := reflect.MakeSlice(arrayType, len(vals), len(vals))
	for i, v := range vals {
		castedValue.Index(i).Set(reflect.ValueOf(v))
	}

	return castedValue.Interface()
}

// restore state updates from mySQL database
func (handler *MySQLSaveStateHandler) RestoreState(ctx *server.EngineCtx, _ string) error {
	gw := ctx.World
	for _, table := range gw.Tables {
		if len(table.EntityToValue) != 0 {
			return fmt.Errorf("table %s is not empty", table.Name)
		}
	}

	for schema, tableAccessor := range handler.schemasToAccessors {
		rows, err := handler.dbConnection.Table(tableNameWithPrefix(tableAccessor.Name(), handler.gameID)).Rows()
		if err != nil {
			return err
		}

		for rows.Next() {
			obj, id, err := convertSQLRowToSchema(rows, schema)
			if err != nil {
				panic(err)
			}

			tableAccessor.Set(gw, id, obj)
		}
	}

	return nil
}

func convertSQLRowToSchema(rows *sql.Rows, schema interface{}) (interface{}, int, error) {
	// Validate that schema is a pointer to a struct
	v := reflect.ValueOf(schema)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, -1, fmt.Errorf("schema must be a pointer to a struct")
	}

	t := v.Elem().Type()
	schemaStruct := reflect.New(t).Elem()
	fieldPointers := make([]interface{}, 0)

	s := stack.New()
	for i := t.NumField() - 1; i >= 0; i-- {
		s.Push(schemaStruct.Field(i))
	}

	idIndex := -1
	for s.Len() != 0 {
		val := s.Pop().(reflect.Value)
		if val.Kind() == reflect.Struct {
			for j := val.NumField() - 1; j >= 0; j-- {
				s.Push(val.Field(j))
			}
		} else {
			fieldPointers = append(fieldPointers, val.Addr().Interface())
		}
	}

	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		if name == "Id" {
			idIndex = i
		}
	}

	err := rows.Scan(fieldPointers...)
	if err != nil {
		return nil, -1, err
	}

	id := schemaStruct.Field(idIndex).Int()

	// Return the populated struct and its primary key
	return schemaStruct.Interface(), int(id), nil
}

func processUpdatesForUpload(tableUpdates []core.TableUpdate) (TableToUpdatesMap, TableToUpdatesMap) {
	// parse the array backwards and store the table updates that are the "latest"
	// ex: if i updated a table row but then deleted it, only the deletion matters
	seenUpdateEntities := make(map[int]bool)
	updates := []core.TableUpdate{}

	for i := len(tableUpdates) - 1; i >= 0; i-- {
		update := tableUpdates[i]
		if !seenUpdateEntities[update.Entity] {
			updates = append(updates, update)
			seenUpdateEntities[update.Entity] = true
		}
	}

	return categorizeTableUpdatesBySchema(updates)

}

// returns: table name -> []value updates

type TableToUpdatesMap map[string][]any

func categorizeTableUpdatesBySchema(updates []core.TableUpdate) (TableToUpdatesMap, TableToUpdatesMap) {
	tableUpdateOperationsByTable := make(TableToUpdatesMap)
	tableRemovalOperationsByTable := make(TableToUpdatesMap)

	for _, update := range updates {
		table := update.Table

		if update.OP == core.UpdateOP {
			tableUpdateOperationsByTable[table] = append(tableUpdateOperationsByTable[table], update.Value)
		} else if update.OP == core.RemovalOP {
			tableRemovalOperationsByTable[table] = append(tableRemovalOperationsByTable[table], update.Value)
		}
	}

	return tableUpdateOperationsByTable, tableRemovalOperationsByTable
}