package server

import (
	"github.com/curio-research/keystone/keystone/ecs"
)

// broadcast state update messages to clients
func BroadcastMessage(w *EngineCtx, worldWithEcsChanges *ecs.GameWorld, jobId int, message string, clientEvents ClientEvents) {
	stateChanges := filterEcsUpdatesWithoutLocal(worldWithEcsChanges.TableUpdates)

	if w.Mode == "debug" {
		if message != "" {
			w.ErrorLog = append(w.ErrorLog, ErrorLog{
				Tick:    w.Ticker.TickNumber,
				Message: message,
			})
		}
	}

	if len(stateChanges) == 0 && message == "" && clientEvents == nil {
		return
	}

	w.Stream.PublishStateChanges(stateChanges, GetJobIdUuid(w.World, jobId), message, clientEvents)
}

func GetJobIdUuid(w *ecs.GameWorld, jobId int) string {
	if jobId == 0 {
		return ""
	}

	job := JobTable.Get(w, jobId)
	return job.TickUuid
}
