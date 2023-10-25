package core

import "math"

func IncludesString(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

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

func SamePos(pos1 Pos, pos2 Pos) bool {
	return pos1.X == pos2.X && pos1.Y == pos2.Y
}

func ArePositionsAdjacent(pos1 Pos, pos2 Pos) bool {
	return SamePos(GetLeftPos(pos1), pos2) || SamePos(GetRightPos(pos1), pos2) || SamePos(GetTopPos(pos1), pos2) || SamePos(GetBottomPos(pos1), pos2)
}

func AreBigTileOriginPositionsAdjacent(pos1 Pos, pos2 Pos, tileWidth int) bool {
	// left, right, top, bottom
	return SamePos(Pos{X: pos1.X - tileWidth, Y: pos1.Y}, pos2) ||
		SamePos(Pos{X: pos1.X + tileWidth, Y: pos1.Y}, pos2) ||
		SamePos(Pos{X: pos1.X, Y: pos1.Y + tileWidth}, pos2) ||
		SamePos(Pos{X: pos1.X, Y: pos1.Y - tileWidth}, pos2)

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
