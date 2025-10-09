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

	fogTexture := rl.LoadRenderTexture(gogue.SCREEN_WIDTH, gogue.SCREEN_HEIGHT)

	enemy := gogue.InitializeEnemy(l)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		camera.Target = player.Position

		mousePos := rl.GetMousePosition()

		mwp := rl.GetScreenToWorld2D(mousePos, camera)

		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.RayWhite)

		l.Draw(mwp, camera)
		enemy.Draw()

		player.Update(dt, *l, mwp)
		enemy.Update(dt, *l, &player)

		if enemy.LockedOntoPlayer(l, &player) {
			enemy.SetupPathfindingTarget(l, &player)
		}

		if player.EnteredDoor(l) {
			l = gogue.GenerateCave(128, 128)
		}

		player.Draw()

		rl.EndMode2D()

		if l.Kind == gogue.LevelKind(gogue.CAVE) {
			rl.BeginTextureMode(fogTexture)
			rl.DrawRectangle(0, 0, gogue.SCREEN_WIDTH, gogue.SCREEN_HEIGHT, rl.Black)
			//playerScreenPos := rl.GetWorldToScreen2D(player.MapPosition.ToVec2(), camera)
			rl.DrawCircleGradient(
				gogue.SCREEN_WIDTH/2,
				gogue.SCREEN_HEIGHT/2,
				500,
				rl.Fade(rl.Blank, 0.5),
				rl.Fade(rl.Black, 1.0),
			)

			rl.EndTextureMode()

			rl.BeginBlendMode(rl.BlendAlpha)
			rl.DrawTextureRec(
				fogTexture.Texture,
				rl.NewRectangle(0, 0, float32(gogue.SCREEN_WIDTH), float32(-gogue.SCREEN_HEIGHT)),
				rl.NewVector2(0, 0),
				rl.White,
			)

			rl.EndBlendMode()
		}

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
