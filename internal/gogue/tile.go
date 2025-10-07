package gogue

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const TILE_SIZE = 16

type TileKind int

const (
	GRASS TileKind = iota
	WALL
	DOOR
)

type Tile struct {
	Position   MapPosition
	Color      rl.Color
	IsPassable bool
	Kind       TileKind
}

func (t *Tile) Draw(mwp rl.Vector2) {

	if t.Kind == TileKind(DOOR) {
		//fmt.Println("DRAWING DOOR")
		// rl.DrawRectangle(int32(t.Position.X-TILE_SIZE/2), int32(t.Position.Y-TILE_SIZE/2), TILE_SIZE*4, TILE_SIZE*4, rl.Red)
		rl.DrawRectangle(int32(t.Position.X*TILE_SIZE), int32(t.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Red)
	} else {
		// fmt.Println("DRAWING NON-DOOR")
		// rl.DrawRectangle(int32(t.Position.X-TILE_SIZE/2), int32(t.Position.Y-TILE_SIZE/2), TILE_SIZE, TILE_SIZE, t.Color)
		rl.DrawRectangle(int32(t.Position.X*TILE_SIZE), int32(t.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, t.Color)
	}

	// rl.DrawRectangle(int32(t.Position.X), int32(t.Position.Y), TILE_SIZE, TILE_SIZE, t.Color)

	if rl.CheckCollisionPointRec(mwp, rl.Rectangle{
		X:      float32(t.Position.X * TILE_SIZE),
		Y:      float32(t.Position.Y * TILE_SIZE),
		Width:  TILE_SIZE,
		Height: TILE_SIZE,
	}) {

		// rl.DrawRectangle(int32(t.Position.X-TILE_SIZE/2), int32(t.Position.Y-TILE_SIZE/2), TILE_SIZE, TILE_SIZE, rl.Yellow)
		rl.DrawRectangle(int32(t.Position.X*TILE_SIZE), int32(t.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Yellow)
	}
}
