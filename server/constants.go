package server

import "time"

type GameMode string

const (
	Dev    GameMode = "dev"
	DevSQL GameMode = "devSQL"
	Prod   GameMode = "prod"

	SaveStateInterval       = time.Second * 10
	DevSQLSaveStateInterval = time.Second

	TickRate = 100
)
