package gogue

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Actor       Actor
	MapPosition MapPosition
	Position    rl.Vector2
	Speed       float32
	MoveTarget  Tile
	CurrentPath []MapPosition
	PathIndex   int
}

type MapPosition struct {
	X int
	Y int
}

func (m MapPosition) ToVec2() rl.Vector2 {
	return rl.Vector2{
		X: (float32(m.X) * TILE_SIZE) + TILE_SIZE/2,
		Y: (float32(m.Y) * TILE_SIZE) + TILE_SIZE/2,
	}
}

func GetMapPositionFromVec(v rl.Vector2) MapPosition {
	return MapPosition{
		X: int(v.X / TILE_SIZE),
		Y: int(v.Y / TILE_SIZE),
	}
}

func InitializePlayer(l *Level) Player {
	return Player{
		// Position:    rl.Vector2{X: 0 + gogue.TILE_SIZE/2, Y: 0 + gogue.TILE_SIZE/2},
		Actor: Actor{
			Stats: InitBaseStats(10, 5, 5, 5, 5),
		},
		Position:    MapPosition{X: l.Width / 2, Y: l.Height / 2}.ToVec2(),
		MapPosition: MapPosition{X: l.Width / 2, Y: l.Height / 2},
		Speed:       100,
		MoveTarget: Tile{
			Position: MapPosition{X: l.Width / 2, Y: l.Height / 2},
		},
	}
}

func (p *Player) IsAtMoveTarget() bool {
	if int(p.Position.X) != p.MoveTarget.Position.X || int(p.Position.Y) != p.MoveTarget.Position.Y {
		return false
	}

	return true
}

func (p *Player) EnteredDoor(l *Level) bool {
	currentTile := l.Get(p.MapPosition.X, p.MapPosition.Y)

	if currentTile.Kind == TileKind(DOOR) {
		fmt.Println("PLAYER HAS ENTERED DOOR")
		return true
	}

	return false
}

func (p *Player) Update(dt float32, l Level, mwp rl.Vector2) {

	p.MapPosition = GetMapPositionFromVec(p.Position)

	pMapPosWorld := p.MapPosition.ToVec2()
	rl.DrawRectangle(int32(pMapPosWorld.X)-TILE_SIZE/2, int32(pMapPosWorld.Y)-TILE_SIZE/2, 16, 16, rl.Orange)

	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		for _, t := range l.Tiles {
			if rl.CheckCollisionPointRec(mwp, rl.Rectangle{
				X:      float32(t.Position.X * TILE_SIZE),
				Y:      float32(t.Position.Y * TILE_SIZE),
				Width:  TILE_SIZE,
				Height: TILE_SIZE,
			}) {
				p.MoveTarget = *t

				//pathToTarget := BreadthFirstSearch(l, p.MapPosition, p.MoveTarget.Position)
				pathToTarget := AStar(l, p.MapPosition, p.MoveTarget.Position)
				p.CurrentPath = pathToTarget
				p.PathIndex = 1
			}
		}
	}

	if !p.IsAtMoveTarget() {

		if len(p.CurrentPath) == 0 {
			return
		}

		// fmt.Println(p.CurrentPath)
		if p.CurrentPath != nil && p.PathIndex < len(p.CurrentPath) {

			rl.DrawCircleV(p.Position, 4, rl.Red)

			for _, step := range p.CurrentPath {
				pos := step.ToVec2()
				rl.DrawRectangle(int32(pos.X)-4, int32(pos.Y)-4, 8, 8, rl.Brown)
			}

			tileTarget := p.CurrentPath[p.PathIndex]
			worldTarget := tileTarget.ToVec2()

			toTarget := rl.Vector2Subtract(worldTarget, p.Position)
			dist := rl.Vector2Length(toTarget)

			step := p.Speed * dt

			if dist <= step {
				p.Position = worldTarget
				p.MapPosition = p.CurrentPath[p.PathIndex]
				p.PathIndex++
				if p.PathIndex >= len(p.CurrentPath) {
					p.CurrentPath = nil
				}
			} else {
				dir := rl.Vector2Normalize(toTarget)
				p.Position = rl.Vector2Add(p.Position, rl.Vector2Scale(dir, p.Speed*dt))
			}
		}

	}
}

func (p *Player) Draw() {
	rl.DrawRectangle(int32(p.Position.X-TILE_SIZE/2), int32(p.Position.Y-TILE_SIZE/2), TILE_SIZE, TILE_SIZE, rl.Blue)
	rl.DrawCircleV(p.Position, 4, rl.Red)
}

func BreadthFirstSearch(l Level, start MapPosition, end MapPosition) []MapPosition {

	// fmt.Println("Start: ", start, "End: ", end)

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

	visited := make(map[MapPosition]bool)
	prev := make(map[MapPosition]MapPosition)

	queue := []MapPosition{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}

		for _, d := range directions {
			neighbour := MapPosition{current.X + d.X, current.Y + d.Y}

			if visited[neighbour] {
				continue
			}

			visited[neighbour] = true
			prev[neighbour] = current

			queue = append(queue, neighbour)

		}
	}

	path := []MapPosition{}
	current := end
	for {
		path = append([]MapPosition{current}, path...)
		if current == start {
			break
		}
		p, ok := prev[current]
		if !ok {
			// fmt.Println("NO PATH FOUND")
			return nil
		}
		current = p
	}

	return path
}

func (m *MapPosition) IsInBounds(boundW int, boundH int) bool {
	if m.X >= boundW || m.X < 0 || m.Y >= boundH || m.Y < 0 {
		return false
	}

	return true
}

func GetLevelTileFromMapPosition(l Level, target MapPosition) Tile {
	for _, lTile := range l.Tiles {
		if lTile.Position.X == target.X && lTile.Position.Y == target.Y {
			return *lTile
		}
	}

	return Tile{}
}

func countVisited(v [][]bool) int {
	count := 0
	for _, row := range v {
		for _, b := range row {
			if b {
				count++
			}
		}
	}

	return count
}
