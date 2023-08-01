package server

import (
	"encoding/json"
	"strconv"

	"github.com/curio-research/go-backend/engine"
	"github.com/curio-research/go-backend/pathfinder"
	"github.com/curio-research/go-backend/server/components"
)

var (
	MoveCalculationTickID = "MoveCalculationTick"
	MoveTickID            = "Move"
	AttackTickID          = "Attack"
	RegenerateTickId      = "Regenerate"
)

// handles troop movement using game ticks
func SingleMoveSystem(ctx *EngineCtx) error {
	jobIds := GetTickJobs(ctx.World, MoveTickID, ctx.Ticker.TickNumber)

	for _, jobId := range jobIds {
		tickDataString, err := components.TickDataStringComponent.Get(ctx.World, jobId)
		if err != nil {
			return err
		}

		var req MovementTickJob
		json.Unmarshal([]byte(tickDataString), &req)

		tilePosition, err := components.PositionComponent.Get(ctx.World, req.TileId)
		if err != nil {
			return err
		}

		w := engine.StartRecordingStateChanges(ctx.World)

		// set troop to new position
		components.PositionComponent.Set(w, req.TroopId, tilePosition)

		w.RemoveEntity(jobId)

		ctx.Stream.PublishStateChanges(w.ExportEcsChanges(), "")
	}

	return nil
}

// handles re-spawning NPCs
func RespawnNPCSystem(ctx *EngineCtx) error {

	npcsInWorld := GetInfantriesOfOwner(ctx.World, 0)
	npcsToSpawn := MaxNPCInWorld - len(npcsInWorld)

	w := engine.StartRecordingStateChanges(ctx.World)

	// deterministically spawn NPCs. TODO: test
	seed := HashNumbers(ctx.RandSeed, ctx.Ticker.TickNumber)
	SpawnNPCsDeterministic(w, npcsToSpawn, seed)

	ctx.Stream.PublishStateChanges(w.ExportEcsChanges(), "")

	return nil
}

// handles attack
func AttackSystem(ctx *EngineCtx) error {
	jobIds := GetTickJobs(ctx.World, AttackTickID, ctx.Ticker.TickNumber)

	for _, jobId := range jobIds {

		w := engine.StartRecordingStateChanges(ctx.World)

		tickDataString, _ := components.TickDataStringComponent.Get(w, jobId)

		// TODO: add error handling and probably stream to logger / client
		var req AttackTickJob
		_ = json.Unmarshal([]byte(tickDataString), &req)

		attackerPosition, _ := components.PositionComponent.Get(w, req.AttackerId)
		targetPosition, _ := components.PositionComponent.Get(w, req.TargetId)

		if engine.WithinDistance(attackerPosition, targetPosition, AttackDistance) {

			attackerHealth, _ := components.HealthComponent.Get(w, req.AttackerId)
			targetHealth, _ := components.HealthComponent.Get(w, req.TargetId)

			// deduct health for target
			if targetHealth < 10 {
				w.RemoveEntity(req.TargetId)
			} else {
				components.HealthComponent.Set(w, req.TargetId, targetHealth-AttackDamage)
			}

			// deduct health for attacker
			if attackerHealth < 10 {
				w.RemoveEntity(req.AttackerId)
			} else {
				components.HealthComponent.Set(w, req.AttackerId, attackerHealth-AttackDamage)
			}

		}

		w.RemoveEntity(jobId)

		ctx.Stream.PublishStateChanges(w.ExportEcsChanges(), "")
	}

	return nil
}

// handles regenerating troops
func RegenerateSystem(ctx *EngineCtx) error {
	jobIds := GetTickJobs(ctx.World, RegenerateTickId, ctx.Ticker.TickNumber)

	for _, jobId := range jobIds {

		dataStr, _ := components.TickDataStringComponent.Get(ctx.World, jobId)

		var req RegenerateTroopsRequest
		_ = json.Unmarshal([]byte(dataStr), &req)

		w := engine.StartRecordingStateChanges(ctx.World)

		troopCount := len(GetInfantriesOfOwner(ctx.World, req.PlayerId))
		troopsToSpawn := MaxInfantryPerPlayer - troopCount

		spawnedTroops := 0

		// loop through all positions from spiral search grid and spawn accordingly
		for _, pos := range SpiralSearchGrid {
			if spawnedTroops < troopsToSpawn {
				searchPos := engine.Pos{X: req.Position.X + pos.X, Y: req.Position.Y + pos.Y}

				if IsPositionInMap(searchPos.X, searchPos.Y, WorldWidth, WorldHeight) {

					entitiesAtPosition := GetEntitiesAtPosition(w, searchPos)

					// if there's only a tile, return this as a valid position
					if len(entitiesAtPosition) == 1 {
						AddInfantry(w, searchPos.X, searchPos.Y, req.PlayerId)
						spawnedTroops++
					}
				}
			}
		}

		ctx.Stream.PublishStateChanges(w.ExportEcsChanges(), "")

		w.RemoveEntity(jobId)

	}

	return nil

}

// handles calculating movement for troops through pathfinding. schedules tick jobs into the future
func MoveCalculationSystem(ctx *EngineCtx) error {
	jobIds := GetTickJobs(ctx.World, MoveCalculationTickID, ctx.Ticker.TickNumber)

	for _, jobId := range jobIds {
		tickDataString, _ := components.TickDataStringComponent.Get(ctx.World, jobId)

		var req SubmitMoveMultiple
		json.Unmarshal([]byte(tickDataString), &req)

		w := engine.StartRecordingStateChanges(ctx.World)

		world2dArray := ConvertLiveMapTo2DArray(ctx.World)
		worldMap := pathfinder.ConstructWorldNew(world2dArray)

		filledPositions := []engine.Pos{}
		centerPos, _ := components.PositionComponent.Get(w, req.TileId)

		// delete all existing queued move jobs
		for _, troopId := range req.TroopIds {
			tickJobsToRemove := GetEntitiesOfTickId(w, strconv.Itoa(troopId))
			RemoveAllEntities(w, tickJobsToRemove)
		}

		// get the first troop position
		for _, troopId := range req.TroopIds {

			// get the troop Position
			startTroopPosition, _ := components.PositionComponent.Get(w, troopId)

			targetPos := FindValidPositionForMove(w, centerPos, filledPositions)

			filledPositions = append(filledPositions, targetPos)

			startPos := pathfinder.Pos{X: startTroopPosition.X, Y: startTroopPosition.Y}
			toPos := pathfinder.Pos{X: targetPos.X, Y: targetPos.Y}

			path, _, _ := pathfinder.AstarPathfinder(startPos, toPos, worldMap)

			// convert to array of position path
			pathPositions := []engine.Pos{}

			// reverse the output path
			for i := len(path) - 1; i >= 0; i-- {
				pos := path[i]
				position, _ := pos.(*pathfinder.Tile)
				pathPositions = append(pathPositions, engine.Pos{X: position.X, Y: position.Y})
			}

			for idx, pos := range pathPositions {

				tileId := GetSmallTileEntityAtPos(w, pos.X, pos.Y)

				// schedule more ticks into the future
				individualMoveRequest := MovementTickJob{
					TroopId: troopId,
					TileId:  tileId,
				}

				tickInMs := (idx + 1) * MoveSpeed

				tick := ctx.Ticker.TickNumber + CalculateTickBasedOnTimeMilliseconds(tickInMs, ctx.Ticker.TickRateMs)

				jsonBytes, _ := json.Marshal(individualMoveRequest)
				jsonString := string(jsonBytes)

				AddTickJob(w, tick, MoveTickID, jsonString, strconv.Itoa(troopId))
			}
		}

		world2dArray = nil
		worldMap = nil

		ctx.Stream.PublishStateChanges(w.ExportEcsChanges(), "")
	}

	return nil
}
