package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/curio-research/go-backend/engine"
	mongoHelper "github.com/curio-research/go-backend/mongo"
	"github.com/curio-research/go-backend/server"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// for debugging using profiler
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	color.HiCyan(asciiArt)
	color.HiCyan("Initiating The Keystone Program ™ ...")
	fmt.Println()

	// initialize a websocket streaming server
	streamServer, err := server.NewStreamServer()
	if err != nil {
		log.Fatal(err)
	}

	randSeed := os.Getenv("RAND_SEED")
	if randSeed == "" {
		log.Fatal("missing RAND_SEED env variable")
	}

	randSeedNumber, err := strconv.Atoi(randSeed)
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	s.Use(server.CORSMiddleware())

	// initialize in-memory game world
	gameWorld := engine.NewGameWorld()
	server.InitializeMap(gameWorld, randSeedNumber)

	// initialize a game tick
	gameTick := server.NewGameTick()

	// this is where you setup the tick schedules for your game
	// different events can happen on different ticks
	// some operations require the game to tick slower
	// ex: health only regenerates every 20 seconds, but you attack every 5 seconds
	tickSchedule := server.NewTickSchedule()

	// TODO: parallelize systems deterministically
	tickSchedule.AddTickSystem(server.NpcRespawnInterval, server.RespawnNPCSystem)
	tickSchedule.AddTickSystem(gameTick.TickRateMs, server.MoveCalculationSystem)
	tickSchedule.AddTickSystem(gameTick.TickRateMs, server.SingleMoveSystem)
	tickSchedule.AddTickSystem(gameTick.TickRateMs, server.AttackSystem)
	tickSchedule.AddTickSystem(gameTick.TickRateMs, server.RegenerateSystem)

	// this is the master game context being passed around, containing pointers to everything
	gameCtx := &server.EngineCtx{
		GameId:              "template-war3-game",
		Paused:              false,
		World:               gameWorld,
		Ticker:              gameTick,
		Stream:              streamServer,
		RandSeed:            randSeedNumber,
		TickTransactionLock: sync.Mutex{},
	}

	gameTick.Setup(gameCtx, tickSchedule)

	// initiate the example mongoDB transaction api interface for DA layer
	mongoDAHandler, err := server.NewMongoDAService()
	if err != nil {
		log.Fatal(err)
	}

	// initiate the example Celestia transaction api interface for DA layer
	// celestiaDAHandler, err := server.NewCelestiaDAService()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// [WIP] example of how you can plug in your DA layer of choice here
	gameCtx.TickTransactionApi = mongoDAHandler

	color.HiMagenta("Tick rate: " + strconv.Itoa(gameTick.TickRateMs) + "ms")

	// in "ghost mode" we keep everything in memory for fast dev
	if os.Getenv("MODE") != "ghost" {

		mongoDB, err := mongoHelper.ConnectToMongoDB()
		if err != nil {
			log.Fatal(err)
		}

		// TODO: refactor this
		server.WorldsCollection = mongoHelper.GetCollection(mongoDB, server.MongoDatabaseName, "worlds")
		server.PlayerCollection = mongoHelper.GetCollection(mongoDB, server.MongoDatabaseName, "player")
	}

	// setup server routes
	server.SetupRoutes(s, gameCtx)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("missing PORT env variable")
	}

	s.Run(":" + port)

}

var asciiArt = `
   ______           _     
  / ____/_  _______(_)___ 
 / /   / / / / ___/ / __ \
/ /___/ /_/ / /  / / /_/ /
\____/\__,_/_/  /_/\____/ 
                           `
