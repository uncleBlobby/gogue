package gogue

import rl "github.com/gen2brain/raylib-go/raylib"

type Level struct {
	Tiles []Tile
}

func (l *Level) Draw(mwp rl.Vector2) {
	for i := 0; i < len(l.Tiles); i++ {
		l.Tiles[i].Draw(mwp)
	}
}
