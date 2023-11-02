package server

import "time"

type GameMode string

const (
	Dev    GameMode = "dev"
	DevSQL GameMode = "devSQL"
	Prod   GameMode = "prod"

	SaveStateInterval       = time.Second * 10
	DevSQLSaveStateInterval = time.Second

	DefaultServerPort    = 9000
	DefaultWebsocketPort = 9001

	DefaultTickRate = 100 // ms

	TickRate = 100
)
