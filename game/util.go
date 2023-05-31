package game

import (
	"fmt"

	"github.com/ethereum/go-ethereum/engine"
)

func SerializePosAsStr(pos engine.Pos) string {
	return fmt.Sprintf("%v,%v", pos.X, pos.Y)
}

// export entire world state as an array of ECS updates
// used by indexer
func ExportWorldState(world *engine.World) engine.ECSUpdateArray {

	allGameData := engine.ECSUpdateArray{}

	// iterate through all components and download ECS data
	for componentName, component := range world.Components {

		// iterate through all entities in component
		for entityId, _ := range component.EntitiesToValue {

			// get the value for the entity
			entityValue := component.EntitiesToValue[entityId]

			// decode bytes based on data type
			decodedValueBasedOnType := engine.DecodeBytesBasedOnDataType(component.DataType, entityValue)

			// add the update to the array
			// value stored as bytes for now
			allGameData = append(allGameData, engine.ECSUpdate{Entity: entityId, Component: componentName, Value: decodedValueBasedOnType})
		}

	}
	return allGameData
}
