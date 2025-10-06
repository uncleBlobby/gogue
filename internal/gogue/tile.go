package gogue

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const TILE_SIZE = 16

type Tile struct {
	Position   MapPosition
	Color      rl.Color
	IsPassable bool
}

func (t *Tile) Draw(mwp rl.Vector2) {
	rl.DrawRectangle(int32(t.Position.X-TILE_SIZE/2), int32(t.Position.Y-TILE_SIZE/2), TILE_SIZE, TILE_SIZE, t.Color)
	// rl.DrawRectangle(int32(t.Position.X), int32(t.Position.Y), TILE_SIZE, TILE_SIZE, t.Color)

	if rl.CheckCollisionPointRec(mwp, rl.Rectangle{
		X:      float32(t.Position.X),
		Y:      float32(t.Position.Y),
		Width:  TILE_SIZE,
		Height: TILE_SIZE,
	}) {

		rl.DrawRectangle(int32(t.Position.X-TILE_SIZE/2), int32(t.Position.Y-TILE_SIZE/2), TILE_SIZE, TILE_SIZE, rl.Yellow)
	}
}
