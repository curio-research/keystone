package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/golang-collections/collections/stack"
	"gorm.io/gorm"
)

type MySQLSaveStateHandler struct {
	dbConnection       *gorm.DB
	gameID             string
	schemasToAccessors map[interface{}]*state.TableBaseAccessor[any]
}

// initialize connection mySQL
func SQLSaveStateHandler(dialector gorm.Dialector, gameID string, schemasToAccessors map[interface{}]*state.TableBaseAccessor[any]) (*MySQLSaveStateHandler, error) {
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
func (m *MySQLSaveStateHandler) initializeDBTables() error {
	db := m.dbConnection
	if db == nil {
		return fmt.Errorf("db connection is nil")
	}

	// all tables that need to be created
	allSchemas := []any{}
	for schema, _ := range m.schemasToAccessors {
		if reflect.TypeOf(schema).Kind() != reflect.Pointer {
			return fmt.Errorf("schema %v is not a pointer to the struct", schema)
		}
		allSchemas = append(allSchemas, schema)
	}

	err := db.AutoMigrate(allSchemas...)
	if err != nil {
		return err
	}

	return nil
}

// save state updates to mySQL database
func (m *MySQLSaveStateHandler) SaveState(tableUpdates []state.TableUpdate) error {
	// process table updates
	tableUpdateOperationsByTable, tableRemovalOperationsByTable := processUpdatesForUpload(tableUpdates)

	// update operations
	for table, updates := range tableUpdateOperationsByTable {
		arr := m.castToSchemaArray(table, updates)
		if arr != nil {
			tx := m.dbConnection.Save(arr)
			if tx.Error != nil {
				return tx.Error
			}
		}
	}

	// removal operations
	for table, removals := range tableRemovalOperationsByTable {
		arr := m.castToSchemaArray(table, removals)
		tx := m.dbConnection.Delete(arr)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

// given a schema type, use the mapping from tables to cast to an array of that type
func (m *MySQLSaveStateHandler) castToSchemaArray(schemaType string, vals []interface{}) interface{} {
	var accessor *state.TableBaseAccessor[any]
	for _, schemaAccessor := range m.schemasToAccessors {
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
func (m *MySQLSaveStateHandler) RestoreState(ctx *server.EngineCtx, _ string) error {
	gw := ctx.World
	for _, table := range gw.Tables {
		if len(table.EntityToValue) != 0 {
			return fmt.Errorf("table %s is not empty", table.Name)
		}
	}

	for schema, tableAccessor := range m.schemasToAccessors {
		rows, err := m.dbConnection.Table(tableNameWithPrefix(tableAccessor.Name(), m.gameID)).Rows()
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
			break
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

func processUpdatesForUpload(tableUpdates []state.TableUpdate) (TableToUpdatesMap, TableToUpdatesMap) {
	// parse the array backwards and store the table updates that are the "latest"
	// ex: if i updated a table row but then deleted it, only the deletion matters
	seenUpdateEntities := make(map[int]bool)
	updates := []state.TableUpdate{}

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

func categorizeTableUpdatesBySchema(updates []state.TableUpdate) (TableToUpdatesMap, TableToUpdatesMap) {
	tableUpdateOperationsByTable := make(TableToUpdatesMap)
	tableRemovalOperationsByTable := make(TableToUpdatesMap)

	for _, update := range updates {
		table := update.Table

		if update.OP == state.UpdateOP {
			tableUpdateOperationsByTable[table] = append(tableUpdateOperationsByTable[table], update.Value)
		} else if update.OP == state.RemovalOP {
			tableRemovalOperationsByTable[table] = append(tableRemovalOperationsByTable[table], update.Value)
		}
	}

	return tableUpdateOperationsByTable, tableRemovalOperationsByTable
}
