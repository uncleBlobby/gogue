package gogue

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const TILE_SIZE = 16

type TileKind int

const (
	GRASS TileKind = iota
	STONE_FLOOR
	WALL
	CAVE_WALL
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
		rl.DrawRectangle(int32(t.Position.X*TILE_SIZE), int32(t.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Red)
	} else if t.Kind == TileKind(STONE_FLOOR) {
		rl.DrawRectangle(int32(t.Position.X*TILE_SIZE), int32(t.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Gray)
	} else if t.Kind == TileKind(CAVE_WALL) {
		rl.DrawRectangle(int32(t.Position.X*TILE_SIZE), int32(t.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.DarkGray)
	} else {
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
