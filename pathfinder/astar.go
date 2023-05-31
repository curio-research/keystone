package astar

import (
	"container/heap"
	"fmt"
	"math/rand"
	"strings"
)

// astar is an A* pathfinding implementation.

// Pather is an interface which allows A* searching on arbitrary objects which
// can represent a weighted graph.
type Pather interface {
	// PathNeighbors returns the direct neighboring nodes of this node which
	// can be pathed to.
	PathNeighbors(cacheLayer World) []Pather
	// PathNeighborCost calculates the exact movement cost to neighbor nodes.
	PathNeighborCost(to Pather) float64
	// PathEstimatedCost is a heuristic method for estimating movement costs
	// between non-adjacent nodes.
	PathEstimatedCost(to Pather) float64
}

// node is a wrapper to store A* data for a Pather node.
type node struct {
	pather Pather
	cost   float64
	rank   float64
	parent *node
	open   bool
	closed bool
	index  int
}

// nodeMap is a collection of nodes keyed by Pather nodes for quick reference.
type nodeMap map[Pather]*node

// get gets the Pather object wrapped in a node, instantiating if required.
func (nm nodeMap) get(p Pather) *node {
	n, ok := nm[p]
	if !ok {
		n = &node{
			pather: p,
		}
		nm[p] = n
	}
	return n
}

// original astar implementation
// Path calculates a short path and the distance between the two Pather nodes.
// If no path is found, found will be false.
func Path(from, to Pather) (path []Pather, distance float64, found bool) {

	startPos := from.(*Tile)
	toPos := to.(*Tile)

	// if the from and to are the same, return nothing
	if startPos == nil || toPos == nil {
		return
	}

	if SamePos(Pos{X: int64(startPos.X), Y: int64(startPos.Y)}, Pos{X: int64(toPos.X), Y: int64(toPos.Y)}) {
		return
	}

	nm := nodeMap{}
	nq := &priorityQueue{}
	heap.Init(nq)
	fromNode := nm.get(from)
	fromNode.open = true
	heap.Push(nq, fromNode)
	for {
		if nq.Len() == 0 {
			// There's no path, return found false.
			return
		}
		current := heap.Pop(nq).(*node)
		current.open = false
		current.closed = true

		if current == nm.get(to) {
			// Found a path to the goal.
			p := []Pather{}
			curr := current
			for curr != nil {
				p = append(p, curr.pather)
				curr = curr.parent
			}

			// remove the first element of path
			// if path is >= 2, remove first element which is the starting position

			if len(p) >= 2 {
				return p[:len(p)-1], current.cost, true
			} else {
				return p, current.cost, true

			}

		}

		for _, neighbor := range current.pather.PathNeighbors(startPos.W) {
			cost := current.cost + current.pather.PathNeighborCost(neighbor)
			neighborNode := nm.get(neighbor)
			if cost < neighborNode.cost {
				if neighborNode.open {
					heap.Remove(nq, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.rank = cost + neighbor.PathEstimatedCost(to)
				neighborNode.parent = current
				heap.Push(nq, neighborNode)
			}
		}
	}
}

// Modified from the original a* package by adding a cache layer which prevents deepCopy
// FIXME: here, the path finding algorithm get's slower as the number of parallel processes run. I suspect it's due to all operations reading the same chunk of memory (?)

func AstarPathfinder(start, end Pos, worldMap World) (path []Pather, distance float64, found bool) {

	cacheLayer := World{}

	cacheLayer.SetFrom(int(start.X), int(start.Y))
	cacheLayer.SetTo(int(end.X), int(end.Y))

	cacheLayer.SetTileWorldRef(int(start.X), int(start.Y), worldMap)
	cacheLayer.SetTileWorldRef(int(end.X), int(end.Y), worldMap)

	from := cacheLayer.From()
	to := cacheLayer.To()

	// if the from and to are the same, return nothing
	startPos := from
	toPos := to

	if startPos == nil || toPos == nil {
		return
	}

	if SamePos(Pos{X: int64(startPos.X), Y: int64(startPos.Y)}, Pos{X: int64(toPos.X), Y: int64(toPos.Y)}) {
		return
	}

	nm := nodeMap{}
	nq := &priorityQueue{}
	heap.Init(nq)
	fromNode := nm.get(from)
	fromNode.open = true
	heap.Push(nq, fromNode)
	for {
		if nq.Len() == 0 {
			// There's no path, return found false.
			return
		}
		current := heap.Pop(nq).(*node)
		current.open = false
		current.closed = true

		if current == nm.get(to) {
			// Found a path to the goal.
			p := []Pather{}
			curr := current
			for curr != nil {
				p = append(p, curr.pather)
				curr = curr.parent
			}

			// remove the first element of path
			// if path is >= 2, remove first element which is the starting position

			if len(p) >= 2 {
				return p[:len(p)-1], current.cost, true
			} else {
				return p, current.cost, true

			}

		}

		for _, neighbor := range current.pather.PathNeighbors(cacheLayer) {
			cost := current.cost + current.pather.PathNeighborCost(neighbor)
			neighborNode := nm.get(neighbor)
			if cost < neighborNode.cost {
				if neighborNode.open {
					heap.Remove(nq, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.rank = cost + neighbor.PathEstimatedCost(to)
				neighborNode.parent = current
				heap.Push(nq, neighborNode)
			}
		}
	}

}

// TODO: make this more efficient
func ConstructMap2dArray(width int, height int, obstacleDensity float64) [][]string {

	world := [][]string{}

	// spawn empty map
	for i := 0; i < height; i++ {
		row := []string{}
		for j := 0; j < width; j++ {
			row = append(row, ".")
		}
		world = append(world, row)
	}

	// select random points to be mountains
	for i := 0; i < int(float64(width)*float64(height)*obstacleDensity); i++ {
		// calculate offset based on f

		x := rand.Intn(width)
		y := rand.Intn(height)

		world[x][y] = "X"

	}

	return world

}

func replaceAtIndex(input string, index int, replacement string) string {
	return input[:index] + string(replacement) + input[index+1:]
}

type Tile struct {
	// Kind is the kind of tile, potentially affecting movement.
	Kind int
	// X and Y are the coordinates of the tile.
	X, Y int
	// W is a reference to the World that the tile is a part of.
	W World
}

type World map[int]map[int]*Tile

var TerrainMap = map[string]int{
	".": KindPlain,
	"~": KindRiver,
	"M": KindMountain,
	"X": KindBlocker,
	"F": KindFrom,
	"T": KindTo,
}

// Kind* constants refer to tile kinds for input and output.
const (
	// KindPlain (.) is a plain tile with a movement cost of 1.
	KindPlain = iota
	// KindRiver (~) is a river tile with a movement cost of 2.
	KindRiver
	// KindMountain (M) is a mountain tile with a movement cost of 3.
	KindMountain
	// KindBlocker (X) is a tile which blocks movement.
	KindBlocker
	// KindFrom (F) is a tile which marks where the path should be calculated
	// from.
	KindFrom
	// KindTo (T) is a tile which marks the goal of the path.
	KindTo
	// KindPath (●) is a tile to represent where the path is in the output.
	KindPath
)

// KindRunes map tile kinds to output runes.
var KindRunes = map[int]rune{
	KindPlain:    '.',
	KindRiver:    '~',
	KindMountain: 'M',
	KindBlocker:  'X',
	KindFrom:     'F',
	KindTo:       'T',
	KindPath:     '●',
}

// RuneKinds map input runes to tile kinds.
var RuneKinds = map[rune]int{
	'.': KindPlain,
	'~': KindRiver,
	'M': KindMountain,
	'X': KindBlocker,
	'F': KindFrom,
	'T': KindTo,
}

// KindCosts map tile kinds to movement costs.
var KindCosts = map[int]float64{
	KindPlain:    1.0,
	KindFrom:     1.0,
	KindTo:       1.0,
	KindRiver:    2.0,
	KindMountain: 3.0,
}

// A Tile is a tile in a grid which implements Pather.

// PathNeighbors returns the neighbors of the tile, excluding blockers and
// tiles off the edge of the board.
func (t *Tile) PathNeighbors(cacheLayer World) []Pather {
	neighbors := []Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {

		newX := t.X + offset[0]
		newY := t.Y + offset[1]

		// see if there's a tile in cache layer first
		if cacheTile := cacheLayer.Tile(newX, newY); cacheTile != nil && cacheTile.Kind != KindBlocker {

			neighbors = append(neighbors, cacheTile)

		} else if n := t.W.Tile(newX, newY); n != nil &&
			n.Kind != KindBlocker {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the directly neighboring tile.
func (t *Tile) PathNeighborCost(to Pather) float64 {
	toT := to.(*Tile)
	return KindCosts[toT.Kind]
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t *Tile) PathEstimatedCost(to Pather) float64 {
	toT := to.(*Tile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

// World is a two dimensional map of Tiles.

// Tile gets the tile at the given coordinates in the world.
func (w World) Tile(x, y int) *Tile {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}

// SetTile sets a tile at the given coordinates in the world.
func (w World) SetTile(t *Tile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*Tile{}
	}
	w[x][y] = t
	t.X = x
	t.Y = y
	t.W = w
}

func (w World) SetTileWorldRef(x, y int, worldRef World) {
	tile := w.Tile(x, y)

	// if tile == nil {
	// 	fmt.Println("tile is nil")
	// 	return
	// }

	tile.W = worldRef
}

// set "from" point
func (w World) SetFrom(x, y int) {
	w.SetTile(&Tile{
		Kind: KindFrom,
	}, int(x), int(y))
}

// set "to" point
func (w World) SetTo(x, y int) {
	w.SetTile(&Tile{
		Kind: KindTo,
	}, int(x), int(y))
}

// FirstOfKind gets the first tile on the board of a kind, used to get the from
// and to tiles as there should only be one of each.
func (w World) FirstOfKind(kind int) *Tile {
	for _, row := range w {
		for _, t := range row {
			if t.Kind == kind {
				return t
			}
		}
	}
	return nil
}

// From gets the from tile from the world.
func (w World) From() *Tile {
	return w.FirstOfKind(KindFrom)
}

// To gets the to tile from the world.
func (w World) To() *Tile {
	return w.FirstOfKind(KindTo)
}

// RenderPath renders a path on top of a world.
func (w World) RenderPath(path []Pather) string {
	width := len(w)
	if width == 0 {
		return ""
	}
	height := len(w[0])
	pathLocs := map[string]bool{}
	for _, p := range path {
		pT := p.(*Tile)
		pathLocs[fmt.Sprintf("%d,%d", pT.X, pT.Y)] = true
	}
	rows := make([]string, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			t := w.Tile(x, y)
			r := ' '
			if pathLocs[fmt.Sprintf("%d,%d", x, y)] {
				r = KindRunes[KindPath]
			} else if t != nil {
				r = KindRunes[t.Kind]
			}
			rows[y] += string(r)
		}
	}
	return strings.Join(rows, "\n")
}

// ParseWorld parses a textual representation of a world into a world map.
func ParseWorld(input string) World {
	w := World{}
	for y, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, raw := range row {
			kind, ok := RuneKinds[raw]
			if !ok {
				kind = KindBlocker
			}
			w.SetTile(&Tile{
				Kind: kind,
			}, x, y)
		}
	}
	return w
}

func ConstructWorldNew(input [][]string) World {
	w := World{}
	for i, row := range input {
		for j, element := range row {
			kind, ok := TerrainMap[element]
			if !ok {
				kind = KindBlocker
			}

			w.SetTile(&Tile{
				Kind: kind,
			}, j, i)
		}
	}

	return w

}

func DeepCopyWorld(input World) World {
	output := World{}

	for i := range input {
		output[i] = make(map[int]*Tile)

		for j := range input[i] {

			output[i][j] = &Tile{
				Kind: input[i][j].Kind,
				X:    i,
				Y:    j,
				W:    output,
			}
		}
	}

	return output
}

func DeepCopy2DArr(input [][]string) [][]string {
	output := make([][]string, len(input))
	for i := range input {
		output[i] = make([]string, len(input[i]))
		copy(output[i], input[i])
	}
	return output
}

// function to check if two positions are the same
func SamePos(pos1 Pos, pos2 Pos) bool {
	return pos1.X == pos2.X && pos1.Y == pos2.Y
}

type Pos struct {
	X int64 `json:"X"`
	Y int64 `json:"Y"`
}
