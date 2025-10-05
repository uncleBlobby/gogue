package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/gogue/internal/gogue"
)

const SCREEN_WIDTH = 800
const SCREEN_HEIGHT = 600

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "gogue 0.01a")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	l := gogue.Level{
		Tiles: []gogue.Tile{},
	}

	for i := -100; i < 100; i++ {
		for j := -100; j < 100; j++ {
			l.Tiles = append(l.Tiles, gogue.Tile{
				Position: rl.Vector2{X: float32(i) * gogue.TILE_SIZE, Y: float32(j) * 16},
				Color:    rl.Green,
			})
		}
	}

	player := gogue.Player{
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

func DrawDebugText(p gogue.Player) {
	rl.DrawText(fmt.Sprintf("PLAYER POS: %.2f, %.2f", p.Position.X, p.Position.Y), 5, 5, 16, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("PLAYER MOVE TARGET: %.2f, %.2f", p.MoveTarget.Position.X, p.MoveTarget.Position.Y), 5, 22, 16, rl.DarkGray)
}
