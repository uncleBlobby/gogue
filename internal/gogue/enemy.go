package gogue

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Enemy struct {
	Actor       Actor
	MapPosition MapPosition
	Position    rl.Vector2
	Speed       float32
	MoveTarget  Tile
	CurrentPath []MapPosition
	PathIndex   int
	Player      *Player
}

func InitializeEnemy(l *Level) Enemy {
	// must make sure the enemy doesn't spawn inside an impassable tile?

	// get rand x in level
	// get rand y in level

	xRand := rand.IntN(l.Width)
	yRand := rand.IntN(l.Height)

	return Enemy{
		Actor: Actor{
			Stats: InitBaseStats(10, 5, 5, 5, 5),
		},
		Position:    MapPosition{X: xRand, Y: yRand}.ToVec2(),
		MapPosition: MapPosition{X: xRand, Y: yRand},
		Speed:       75,
		MoveTarget: Tile{
			Position: MapPosition{X: xRand, Y: yRand},
		},
	}
}

func (p *Enemy) IsAtMoveTarget() bool {
	if int(p.Position.X) != p.MoveTarget.Position.X || int(p.Position.Y) != p.MoveTarget.Position.Y {
		return false
	}

	return true
}

func (e *Enemy) Draw() {
	rl.DrawRectangle(int32(e.Position.X-TILE_SIZE/2), int32(e.Position.Y-TILE_SIZE/2), TILE_SIZE, TILE_SIZE, rl.Red)
}

func (e *Enemy) Update(dt float32, l Level) {
	e.MapPosition = GetMapPositionFromVec(e.Position)

	// pMapPosWorld := e.MapPosition.ToVec2()

	if !e.IsAtMoveTarget() {

		if len(e.CurrentPath) == 0 {
			return
		}

		// fmt.Println(e.CurrentPath)
		if e.CurrentPath != nil && e.PathIndex < len(e.CurrentPath) {

			rl.DrawCircleV(e.Position, 4, rl.Red)

			for _, step := range e.CurrentPath {
				pos := step.ToVec2()
				rl.DrawRectangle(int32(pos.X)-4, int32(pos.Y)-4, 8, 8, rl.Brown)
			}

			tileTarget := e.CurrentPath[e.PathIndex]
			worldTarget := tileTarget.ToVec2()

			toTarget := rl.Vector2Subtract(worldTarget, e.Position)
			dist := rl.Vector2Length(toTarget)

			step := e.Speed * dt

			if dist <= step {
				e.Position = worldTarget
				e.MapPosition = e.CurrentPath[e.PathIndex]
				e.PathIndex++
				if e.PathIndex >= len(e.CurrentPath) {
					e.CurrentPath = nil
				}
			} else {
				dir := rl.Vector2Normalize(toTarget)
				e.Position = rl.Vector2Add(e.Position, rl.Vector2Scale(dir, e.Speed*dt))
			}
		}

	}
}
