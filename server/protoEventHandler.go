package server

import (
	"strings"

	serverpb "github.com/curio-research/keystone/server/serverpb/output"
	"github.com/curio-research/keystone/state"
)

// Protobuf based implementation for our game

type ProtoBasedErrorHandler struct {
}

// format message into protobuf
func (h *ProtoBasedErrorHandler) FormatMessage(transactionUuidIdentifier int, errorMessage string) *NetworkMessage {

	msg, _ := NewMessage(0, uint32(serverpb.CMD_S2C_Error), uint32(transactionUuidIdentifier), &serverpb.S2C_ErrorMessage{
		Content: errorMessage,
	})

	return msg
}

type ProtoBasedBroadcastHandler struct {
}

func (h *ProtoBasedBroadcastHandler) BroadcastMessage(ctx *EngineCtx, clientEvents []ClientEvent) {
	stateChanges := filterTableUpdatesWithoutLocal(ctx.World.TableUpdates)

	if ctx.ShouldRecordError {
		// loop through all client client events and see which one is an error
		for _, clientEvent := range clientEvents {
			if clientEvent.NetworkMessage.GetCommand() == uint32(serverpb.CMD_S2C_Error) {

				// decode the error message string from serverpb and log it
				data := &serverpb.S2C_ErrorMessage{}
				clientEvent.NetworkMessage.GetProtoMessage(data)

				ctx.ErrorLog = append(ctx.ErrorLog, ErrorLog{
					Tick:    ctx.GameTick.TickNumber,
					Message: data.Content,
				})
			}
		}
	}

	if len(stateChanges) == 0 && clientEvents == nil {
		return
	}

	ctx.Stream.PublishStateChanges(stateChanges, clientEvents)
}

func filterTableUpdatesWithoutLocal(tableUpdates state.TableUpdateArray) state.TableUpdateArray {
	// if the table name starts with the word local, then filter it out and not broadcast it

	filteredUpdates := state.TableUpdateArray{}

	for _, tableUpdate := range tableUpdates {
		if !strings.HasPrefix(tableUpdate.Table, "local") {
			filteredUpdates = append(filteredUpdates, tableUpdate)
		}
	}

	return filteredUpdates
}
