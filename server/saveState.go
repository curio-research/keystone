package server

import (
	"time"
)

// TODO: get back get state route
type DownloadWorldRequest struct {
	Components []string `json:"components"`
}

// general interface to implement
type ISaveDatabase interface {
	// save game updates to database
	SaveUpdatesToDB(ctx *EngineCtx)

	RestoreStateFromDB(ctx *EngineCtx, gameId string)
}

// game loop that triggers the save world state
// set up a game tick
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

					if tick%saveInterval == 0 {
						if saveInterval != 0 {
							ctx.SaveDatabase.SaveUpdatesToDB(ctx)
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
