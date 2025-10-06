package gogue

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	Tiles  []Tile
	Width  int
	Height int
}

func (l *Level) Draw(mwp rl.Vector2, camera rl.Camera2D) {

	visibleCols := SCREEN_WIDTH / TILE_SIZE
	visibleRows := SCREEN_HEIGHT / TILE_SIZE

	startX := int(camera.Target.X/TILE_SIZE) - visibleCols/2 - 1
	startY := int(camera.Target.Y/TILE_SIZE) - visibleRows/2 - 1
	endX := startX + visibleCols + 2
	endY := startY + visibleCols + 2

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			if tile := l.Get(x, y); tile != nil {
				tile.Draw(mwp)
			}
		}
	}

	// for i := 0; i < len(l.Tiles); i++ {
	// 	l.Tiles[i].Draw(mwp)
	// }
}

func (l *Level) Index(x, y int) int {
	return y*l.Width + x
}

func (l *Level) Get(x, y int) *Tile {
	if x < 0 || y < 0 || x >= l.Width || y >= l.Height {
		return nil
	}

	return &l.Tiles[l.Index(x, y)]
}

func (l *Level) IsWalkable(pos MapPosition) bool {
	t := l.Get(pos.X, pos.Y)
	if t == nil {
		//panic("tile is nil")
		//fmt.Println("TILE IS NIL")
	}
	// fmt.Println("TILE IS PASSABLE: ", t.IsPassable)
	return t != nil && t.IsPassable
}
