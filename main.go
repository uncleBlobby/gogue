package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/gogue/internal/gogue"
)

func main() {
	rl.InitWindow(gogue.SCREEN_WIDTH, gogue.SCREEN_HEIGHT, "gogue 0.01a")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	l := gogue.GenerateLevel(128, 128)

	player := gogue.InitializePlayer(l)

	camera := rl.Camera2D{
		Target: player.Position,
		Offset: rl.Vector2{X: gogue.SCREEN_WIDTH / 2, Y: gogue.SCREEN_HEIGHT / 2},
		Zoom:   1.0,
	}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		camera.Target = player.Position

		mousePos := rl.GetMousePosition()

		mwp := rl.GetScreenToWorld2D(mousePos, camera)

		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.RayWhite)

		l.Draw(mwp, camera)

		player.Update(dt, *l, mwp)

		if player.EnteredDoor(l) {
			l = gogue.GenerateCave(128, 128)
		}

		player.Draw()

		rl.EndMode2D()
		rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), 0, 0, 24, rl.DarkGray)

		//		DrawDebugText(player)

		rl.EndDrawing()
	}
}

func DrawDebugText(p gogue.Player) {
	// rl.DrawText(fmt.Sprintf("PLAYER POS: %d, %d", p.Position.X, p.Position.Y), 5, 5, 16, rl.DarkGray)
	// rl.DrawText(fmt.Sprintf("PLAYER MOVE TARGET: %d, %d", p.MoveTarget.Position.X, p.MoveTarget.Position.Y), 5, 22, 16, rl.DarkGray)
	// rl.DrawText(fmt.Sprintf("AT MOVE TARGET: %t", p.IsAtMoveTarget()), 5, 40, 16, rl.DarkGray)

	rl.DrawText(fmt.Sprintf("p.MapPosition: %d, %d", p.MapPosition.X, p.MapPosition.Y), 250, 10, 16, rl.Black)
	// rl.DrawText(fmt.Sprintf("p.MoveTarget.Position: %d, %d", p.MoveTarget.Position.X, p.MoveTarget.Position.Y), 250, 30, 16, rl.Black)
}
