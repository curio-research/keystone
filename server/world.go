package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/curio-research/go-backend/engine"
	"github.com/curio-research/go-backend/server/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveWorld(ctx *EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		_ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		w := ctx.World

		worldId := uuid.New().String()

		gameWorld := models.GameWorld{}
		defer cancel()

		gameWorld.SnapshotId = worldId
		gameWorld.Entities = w.Entities.GetAll()
		gameWorld.EntitiesNonce = w.EntitiesNonce
		gameWorld.Components = []models.Component{}

		for componentName, component := range w.Components {
			componentToSave := models.Component{}

			componentToSave.Name = componentName
			componentToSave.DataType = int(component.DataType)
			componentToSave.Entities = component.Entities.GetAll()
			componentToSave.EntitiesToValue = component.EntitiesToValue

			valueToEntitiesToSave := map[string][]int{}

			for rawVal, entities := range component.ValueToEntities {
				// encode value to string
				stringVal := engine.EncodeToStringBasedOnDataType(component.DataType, rawVal)
				valueToEntitiesToSave[stringVal] = entities.GetAll()

			}

			componentToSave.ValueToEntities = valueToEntitiesToSave

			// add to total array
			gameWorld.Components = append(gameWorld.Components, componentToSave)
		}

		result, err := WorldsCollection.InsertOne(_ctx, gameWorld)
		if err != nil {
			c.JSON(http.StatusInternalServerError, GeneralResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, GeneralResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})

	}
}

func FetchLoadWorld(ctx *EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		_ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		snapshotId := c.DefaultQuery("snapshotId", "")
		shouldLoadIntoMemory, _ := strconv.ParseBool(c.DefaultQuery("load", "false"))

		var fetchedWorld models.GameWorld
		defer cancel()

		err := WorldsCollection.FindOne(_ctx, bson.M{"snapshotid": snapshotId}).Decode(&fetchedWorld)

		if err != nil {
			c.JSON(http.StatusInternalServerError, GeneralResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if shouldLoadIntoMemory {
			// load game into memory

			// clear game world first
			newWorld := engine.NewGameWorld()

			newWorld.Entities = engine.ArrayToSparseSet(fetchedWorld.Entities)
			newWorld.EntitiesNonce = fetchedWorld.EntitiesNonce

			for _, component := range fetchedWorld.Components {
				newComponentToSave := engine.Component{}

				newComponentToSave.DataType = engine.DataType(component.DataType)

				// TODO: include in saved components in DB?
				newComponentToSave.ShouldStoreValueToEntities = true
				newComponentToSave.Entities = engine.ArrayToSparseSet(component.Entities)

				// save entity to component pair
				newComponentToSave.EntitiesToValue = map[int]string{}

				parsedEntitiesToValueKvPairs, _ := ParseKVPair(component.EntitiesToValue)
				for _, kvPair := range parsedEntitiesToValueKvPairs {
					entity, _ := strconv.Atoi(kvPair.Entity)
					newComponentToSave.EntitiesToValue[entity] = kvPair.StringVal
				}

				newComponentToSave.ValueToEntities = map[any]*engine.SparseSet{}

				parsedValueToEntitiesKvPairs, _ := ParseAnyToArrayKvPair(component.ValueToEntities)

				for _, kvPair := range parsedValueToEntitiesKvPairs {
					newComponentToSave.ValueToEntities[kvPair.Data] = engine.ArrayToSparseSet(kvPair.Entities)
				}

				// set component value to newly created world struct
				newWorld.Components[component.Name] = newComponentToSave
			}

			// deep copy and set into memory
			// TODO: test this
			*ctx.World = newWorld.DeepCopy()

		}

		c.JSON(http.StatusOK, GeneralResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": fetchedWorld}})

	}

}

// export entire ECS world
func DownloadWorld(ctx *EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req DownloadWorldRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		exportObj := engine.ExportECSWorld{}

		for _, componentName := range req.Components {
			component, exists := ctx.World.Components[componentName]

			if exists {
				componentForExport := engine.ExportComponent{}
				componentForExport.EntitiesToValue = make(map[int]any)
				componentForExport.ValueToEntities = make(map[string][]int)

				componentForExport.DataType = component.DataType

				// loop through value to entity
				for data, entities := range component.ValueToEntities {
					if entities != nil {
						encodedValToString := engine.EncodeToStringBasedOnDataType(component.DataType, data)

						componentForExport.ValueToEntities[encodedValToString] = entities.GetAll()
					}
				}

				// put this in a big struct
				exportObj[componentName] = componentForExport
			}
		}

		returnObj := struct {
			Status string `json:"status"`
			Data   any    `json:"data"`
		}{
			Status: "success",
			Data:   exportObj,
		}

		c.JSON(http.StatusOK, returnObj)
	}
}

type EntityStringValueKvPair struct {
	Entity    string `json:"Key"`
	StringVal string `json:"Value"`
}

type AnyToArrayKvPair struct {
	Data     any   `json:"Key"`
	Entities []int `json:"Value"`
}

func ParseKVPair(data interface{}) ([]EntityStringValueKvPair, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var entities []EntityStringValueKvPair
	err = json.Unmarshal(bytes, &entities)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func ParseAnyToArrayKvPair(data interface{}) ([]AnyToArrayKvPair, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var entities []AnyToArrayKvPair
	err = json.Unmarshal(bytes, &entities)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

type GeneralResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
