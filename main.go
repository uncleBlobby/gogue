package main

import (
	"fmt"
	"math/rand"

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
		Tiles:  []gogue.Tile{},
		Width:  100,
		Height: 100,
	}

	for j := 0; j < l.Height; j++ {
		for i := 0; i < l.Width; i++ {

			if j == 10 {
				l.Tiles = append(l.Tiles, gogue.Tile{
					Position:   gogue.MapPosition{X: (i)*gogue.TILE_SIZE + gogue.TILE_SIZE/2, Y: (j)*16 + gogue.TILE_SIZE/2},
					Color:      rl.Gray,
					IsPassable: false,
				})
			} else if rand.Float32() < 0.95 {
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

	fmt.Println("TILE 10, 10 isPassable: ", l.IsWalkable(gogue.MapPosition{10, 10}))

	player := gogue.Player{
		// Position:    rl.Vector2{X: 0 + gogue.TILE_SIZE/2, Y: 0 + gogue.TILE_SIZE/2},
		Position:    gogue.MapPosition{X: 25, Y: 25}.ToVec2(),
		MapPosition: gogue.MapPosition{X: 25, Y: 25},
		Speed:       100,
		MoveTarget: gogue.Tile{
			Position: gogue.MapPosition{X: 25, Y: 25},
		},
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

		// mouseTilePosition := gogue.GetMapPositionFromVec(mwp)

		// player.Update(dt, l, mwp)

		rl.BeginDrawing()

		rl.BeginMode2D(camera)

		rl.ClearBackground(rl.RayWhite)

		l.Draw(mwp)
		player.Draw()

		player.Update(dt, l, mwp)

		rl.EndMode2D()

		DrawDebugText(player)
		//rl.DrawText(fmt.Sprintf("MouseTILE: %d, %d", mouseTilePosition.X, mouseTilePosition.Y), 25, 25, 36, rl.Black)

		//tileUnderMouse := l.Get(mouseTilePosition.X, mouseTilePosition.Y)

		//rl.DrawText(fmt.Sprintf("isWalkable: %t", tileUnderMouse.IsPassable), 50, 300, 32, rl.Black)

		// for _, t := range l.Tiles {
		// 	if !t.IsPassable {

		// 		tWorldPos := t.Position
		// 		// fmt.Println("NON WALKABLE: ", tWorldPos)

		// 		screenY := tWorldPos.Y - (7 * gogue.TILE_SIZE)
		// 		screenX := tWorldPos.X - (1 * gogue.TILE_SIZE)

		// 		rl.DrawCircle(int32(screenX+gogue.TILE_SIZE/2), int32(screenY+gogue.TILE_SIZE/2), 3, rl.Red)
		// 	}
		// }

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
