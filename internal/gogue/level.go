package gogue

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	Tiles  []Tile
	Width  int
	Height int
}

func (l *Level) Draw(mwp rl.Vector2) {
	for i := 0; i < len(l.Tiles); i++ {
		l.Tiles[i].Draw(mwp)
	}
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
