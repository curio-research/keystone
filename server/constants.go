package server

type GameMode string

const (
	Dev    GameMode = "dev"
	DevSQL GameMode = "devSQL"
	Prod   GameMode = "prod"

	SaveStateInterval       = 10
	DevSQLSaveStateInterval = 1

	TickRate = 100
)
