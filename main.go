package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const SCREEN_WIDTH = 800
const SCREEN_HEIGHT = 600

type Level struct {
	Tiles []Tile
}

type Tile struct {
	Position rl.Vector2
	Color    rl.Color
}

type Player struct {
	Position   rl.Vector2
	Speed      float32
	MoveTarget Tile
}

func (p *Player) IsAtMoveTarget() bool {
	if p.Position.X != p.MoveTarget.Position.X || p.Position.Y != p.MoveTarget.Position.Y {
		return false
	}

	return true
}

func (p *Player) Update(dt float32, l Level, mwp rl.Vector2) {

	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		for _, t := range l.Tiles {
			if rl.CheckCollisionPointRec(mwp, rl.Rectangle{
				X:      t.Position.X,
				Y:      t.Position.Y,
				Width:  16,
				Height: 16,
			}) {
				p.MoveTarget = t
			}
		}
	}

	if !p.IsAtMoveTarget() {

		moveDir := rl.Vector2Subtract(p.MoveTarget.Position, p.Position)
		// move toward target

		p.Position.X += moveDir.X * p.Speed * dt
		p.Position.Y += moveDir.Y * p.Speed * dt
		// find path
	}
}

func (p *Player) Draw() {
	rl.DrawRectangle(int32(p.Position.X), int32(p.Position.Y), 16, 16, rl.Blue)
}

func (l *Level) Draw(mwp rl.Vector2) {
	for i := 0; i < len(l.Tiles); i++ {
		l.Tiles[i].Draw(mwp)
	}
}

func (t *Tile) Draw(mwp rl.Vector2) {
	rl.DrawRectangle(int32(t.Position.X), int32(t.Position.Y), 16, 16, t.Color)

	if rl.CheckCollisionPointRec(mwp, rl.Rectangle{
		X:      t.Position.X,
		Y:      t.Position.Y,
		Width:  16,
		Height: 16,
	}) {
		rl.DrawRectangle(int32(t.Position.X), int32(t.Position.Y), 16, 16, rl.Yellow)
	}
}

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "gogue 0.01a")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	l := Level{
		Tiles: []Tile{},
	}

	for i := -100; i < 100; i++ {
		for j := -100; j < 100; j++ {
			l.Tiles = append(l.Tiles, Tile{
				Position: rl.Vector2{X: float32(i) * 16, Y: float32(j) * 16},
				Color:    rl.Green,
			})
		}
	}

	player := Player{
		Position: rl.Vector2{X: 0, Y: 0},
		Speed:    1,
	}

	camera := rl.Camera2D{
		Target: player.Position,
		Offset: rl.Vector2{X: SCREEN_WIDTH / 2, Y: SCREEN_HEIGHT / 2},
		Zoom:   1.0,
	}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		mousePos := rl.GetMousePosition()
		mwp := rl.GetScreenToWorld2D(mousePos, camera)

		player.Update(dt, l, mwp)

		rl.BeginDrawing()

		rl.BeginMode2D(camera)

		rl.ClearBackground(rl.RayWhite)

		l.Draw(mwp)
		player.Draw()

		rl.EndMode2D()

		DrawDebugText(player)

		rl.EndDrawing()

	}
}

func DrawDebugText(p Player) {
	rl.DrawText(fmt.Sprintf("PLAYER POS: %.2f, %.2f", p.Position.X, p.Position.Y), 5, 5, 16, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("PLAYER MOVE TARGET: %.2f, %.2f", p.MoveTarget.Position.X, p.MoveTarget.Position.Y), 5, 22, 16, rl.DarkGray)
}
