package gogue

import (
	"math"
)

type Node struct {
	Pos    MapPosition
	G, H   float64
	F      float64
	Parent *Node
}

func heuristic(a, b MapPosition) float64 {
	dx := math.Abs(float64(a.X - b.X))
	dy := math.Abs(float64(a.Y - b.Y))
	return math.Max(dx, dy)
}

func AStar(level Level, start, goal MapPosition) []MapPosition {

	goal = MapPosition{goal.X, goal.Y}

	var directions = []MapPosition{
		{0, 1}, // down

		{1, 0}, // right

		{0, -1}, // up

		{-1, 0}, // left

		{1, 1}, // down right

		{1, -1}, // up right

		{-1, -1}, // up left

		{-1, 1}, // down left
	}

	open := map[MapPosition]*Node{}
	closed := map[MapPosition]bool{}

	startNode := &Node{Pos: start, G: 0, H: heuristic(start, goal)}
	startNode.F = startNode.G + startNode.H
	open[start] = startNode

	var current *Node

	for len(open) > 0 {
		current = nil
		for _, n := range open {
			if current == nil || n.F < current.F {
				current = n
			}
		}

		if current.Pos == goal {
			var path []MapPosition
			for n := current; n != nil; n = n.Parent {
				path = append([]MapPosition{n.Pos}, path...)
			}
			return path
		}

		delete(open, current.Pos)
		closed[current.Pos] = true

		for _, dir := range directions {
			neighbourPos := MapPosition{current.Pos.X + dir.X, current.Pos.Y + dir.Y}

			if dir.X != 0 && dir.Y != 0 {
				if !level.IsWalkable(MapPosition{current.Pos.X + dir.X, current.Pos.Y}) || !level.IsWalkable(MapPosition{current.Pos.X, current.Pos.Y + dir.Y}) {
					continue
				}
			}

			if !neighbourPos.IsInBounds(level.Width, level.Height) {
				continue
			}

			if !level.IsWalkable(neighbourPos) {
				continue
			}

			if closed[neighbourPos] {
				continue
			}

			tentativeG := current.G + 1 // or sqrt2 for diagonals?
			if dir.X != 0 && dir.Y != 0 {
				tentativeG = current.G + 1.41421356237
			}

			neighbour, exists := open[neighbourPos]
			if !exists || tentativeG < neighbour.G {
				if !exists {
					neighbour = &Node{Pos: neighbourPos}
				}
				neighbour.G = tentativeG
				neighbour.H = heuristic(neighbourPos, goal)
				neighbour.F = neighbour.G + neighbour.H
				neighbour.Parent = current
				open[neighbourPos] = neighbour
			}
		}
	}

	return nil
}
