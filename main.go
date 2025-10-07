package main

import (
	"fmt"
	"math/rand"

	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/gogue/internal/gogue"
)

func main() {
	rl.InitWindow(gogue.SCREEN_WIDTH, gogue.SCREEN_HEIGHT, "gogue 0.01a")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	//rl.SetConfigFlags(rl.FlagVsyncHint)

	l := gogue.Level{
		Tiles:  []gogue.Tile{},
		Width:  256,
		Height: 256,
	}

	alpha := 2.0
	beta := 2.0
	n := int32(3)
	seed := rand.Int63()

	p := perlin.NewPerlin(alpha, beta, n, seed)

	scale := 0.1

	for j := 0; j < l.Height; j++ {
		for i := 0; i < l.Width; i++ {

			val := p.Noise2D(float64(i)*scale, float64(j)*scale)
			// val := 0.6*p.Noise2D(float64(i)*0.05, float64(j)*0.05) + 0.75*p.Noise2D(float64(i)*0.15, float64(j)*0.15)

			val = (val + 1) / 2

			if j == 10 {
				l.Tiles = append(l.Tiles, gogue.Tile{
					Position:   gogue.MapPosition{X: (i)*gogue.TILE_SIZE + gogue.TILE_SIZE/2, Y: (j)*16 + gogue.TILE_SIZE/2},
					Color:      rl.Gray,
					IsPassable: false,
				})
			} else if val > 0.4 {
				l.Tiles = append(l.Tiles, gogue.Tile{
					Position:   gogue.MapPosition{X: (i)*gogue.TILE_SIZE + gogue.TILE_SIZE/2, Y: (j)*16 + gogue.TILE_SIZE/2},
					Color:      rl.Green,
					IsPassable: true,
				})
			} else {
				l.Tiles = append(l.Tiles, gogue.Tile{
					Position:   gogue.MapPosition{X: (i)*gogue.TILE_SIZE + gogue.TILE_SIZE/2, Y: (j)*16 + gogue.TILE_SIZE/2},
					Color:      rl.Gray,
					IsPassable: false,
				})
			}
		}
	}

	// fmt.Println("TILE 10, 10 isPassable: ", l.IsWalkable(gogue.MapPosition{10, 10}))

	player := gogue.Player{
		// Position:    rl.Vector2{X: 0 + gogue.TILE_SIZE/2, Y: 0 + gogue.TILE_SIZE/2},
		Position:    gogue.MapPosition{X: l.Width / 2, Y: l.Height / 2}.ToVec2(),
		MapPosition: gogue.MapPosition{X: l.Width / 2, Y: l.Height / 2},
		Speed:       100,
		MoveTarget: gogue.Tile{
			Position: gogue.MapPosition{X: l.Width / 2, Y: l.Height / 2},
		},
	}

	camera := rl.Camera2D{
		Target: player.Position,
		Offset: rl.Vector2{X: gogue.SCREEN_WIDTH / 2, Y: gogue.SCREEN_HEIGHT / 2},
		Zoom:   1.0,
	}

	l.InitRenderTexture()

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		camera.Target = player.Position

		mousePos := rl.GetMousePosition()

		mwp := rl.GetScreenToWorld2D(mousePos, camera)

		l.RedrawStaticLayer(mwp)

		// mouseTilePosition := gogue.GetMapPositionFromVec(mwp)

		// player.Update(dt, l, mwp)

		rl.BeginDrawing()

		rl.BeginMode2D(camera)

		rl.ClearBackground(rl.RayWhite)

		//l.Draw(mwp, camera)
		l.DrawRenderTexture(mwp, camera)

		player.Update(dt, l, mwp)

		player.Draw()

		rl.EndMode2D()
		rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), 0, 0, 24, rl.DarkGray)

		// DrawDebugText(player)

		rl.EndDrawing()
	}
}

func DrawDebugText(p gogue.Player) {
	rl.DrawText(fmt.Sprintf("PLAYER POS: %d, %d", p.Position.X, p.Position.Y), 5, 5, 16, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("PLAYER MOVE TARGET: %d, %d", p.MoveTarget.Position.X, p.MoveTarget.Position.Y), 5, 22, 16, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("AT MOVE TARGET: %t", p.IsAtMoveTarget()), 5, 40, 16, rl.DarkGray)

	rl.DrawText(fmt.Sprintf("p.MapPosition: %d, %d", p.MapPosition.X, p.MapPosition.Y), 250, 10, 16, rl.Black)
	rl.DrawText(fmt.Sprintf("p.MoveTarget.Position: %d, %d", p.MoveTarget.Position.X, p.MoveTarget.Position.Y), 250, 30, 16, rl.Black)
}
