package engine

import "math"

func GetLeftPos(pos Pos) Pos {
	return Pos{X: pos.X - 1, Y: pos.Y}
}

func GetRightPos(pos Pos) Pos {
	return Pos{X: pos.X + 1, Y: pos.Y}
}

func GetTopPos(pos Pos) Pos {
	return Pos{X: pos.X, Y: pos.Y - 1}
}

func GetBottomPos(pos Pos) Pos {
	return Pos{X: pos.X, Y: pos.Y + 1}
}

func WithinDistance(pos1 Pos, pos2 Pos, distance int) bool {
	return math.Abs(float64(pos1.X-pos2.X)) <= float64(distance) && math.Abs(float64(pos1.Y-pos2.Y)) <= float64(distance)
}

func ContainsPositions(posList []Pos, pos Pos) bool {
	for _, p := range posList {
		if p.X == pos.X && p.Y == pos.Y {
			return true
		}
	}

	return false
}

func (world *World) DeepCopy() World {
	newWorld := World{}

	newWorld.Parent = world.Parent
	newWorld.SnapshotMode = world.SnapshotMode

	newWorld.Entities = world.Entities.DeepCopy()
	newWorld.EntitiesNonce = world.EntitiesNonce

	newWorld.Components = make(map[string]Component)

	// TODO: should be a way to deep copy in parallel
	for componentName, component := range world.Components {
		newWorld.Components[componentName] = component.DeepCopy()
	}

	return newWorld
}
