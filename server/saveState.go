package server

import (
	"time"
)

// TODO: add back get state route
type DownloadWorldRequest struct {
	Components []string `json:"components"`
}

// general interface to implement
type ISaveState interface {
	// save game updates to database
	SaveState(ctx *EngineCtx)

	RestoreState(ctx *EngineCtx, gameId string)
}

// game loop that triggers the save world state
func SetupSaveStateLoop(ctx *EngineCtx, saveInterval int) {
	tickerTime := time.Duration(1_000) * time.Millisecond
	ticker := time.NewTicker(tickerTime)

	quit := make(chan struct{})

	tick := 0 // seconds

	go func() {
		for {
			select {
			case <-ticker.C:

				if ctx.IsLive {
					tick++

					if saveInterval != 0 {
						if tick%saveInterval == 0 {
							ctx.SaveStateHandler.SaveState(ctx)
						}
					}

				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
