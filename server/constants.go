package server

import "time"

type GameMode string

const (
	Dev       GameMode = "dev"
	DevMySQL  GameMode = "devMySQL"
	DevSQLite GameMode = "devSQLite"
	Prod      GameMode = "prod"

	SaveStateInterval    = time.Second * 10
	DevSaveStateInterval = time.Millisecond * 100

	DefaultServerPort    = 9000
	DefaultWebsocketPort = 9001

	TickRate = 100
)
