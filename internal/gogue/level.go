package gogue

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	Tiles     []Tile
	Width     int
	Height    int
	RenderTex rl.RenderTexture2D
	Dirty     bool
}

func (l *Level) InitRenderTexture() {
	l.RenderTex = rl.LoadRenderTexture(int32(l.Width*TILE_SIZE), int32(l.Height*TILE_SIZE))
	l.Dirty = true
}

func (l *Level) RedrawStaticLayer(mwp rl.Vector2) {
	if !l.Dirty {
		return
	}

	rl.BeginTextureMode(l.RenderTex)
	rl.ClearBackground(rl.Blank)

	rl.DisableDepthTest()
	rl.DisableBackfaceCulling()

	for y := 0; y < l.Height; y++ {
		for x := 0; x < l.Width; x++ {
			t := l.Get(x, y)
			if t != nil {
				t.Draw(mwp)
			}
		}
	}

	rl.EndTextureMode()
	l.Dirty = false
}

func (l *Level) DrawRenderTexture(mwp rl.Vector2, camera rl.Camera2D) {
	//rl.BeginMode2D(camera)
	rl.DrawTextureRec(
		l.RenderTex.Texture,
		rl.Rectangle{0, 0, float32(l.Width * TILE_SIZE), -float32(l.Height * TILE_SIZE)},
		rl.Vector2{0, 0},
		rl.White,
	)
	//rl.EndMode2D()
}

func (l *Level) PlainDraw(mwp rl.Vector2, camera rl.Camera2D) {

	tileCount := 0

	visibleCols := SCREEN_WIDTH / TILE_SIZE
	visibleRows := SCREEN_HEIGHT / TILE_SIZE

	startX := int(camera.Target.X/TILE_SIZE) - visibleCols/2 - 1
	startY := int(camera.Target.Y/TILE_SIZE) - visibleRows/2 - 1
	endX := startX + visibleCols + 2
	endY := startY + visibleRows + 2

	if startX < 0 {
		startX = 0
	}

	if startY < 0 {
		startY = 0
	}

	if endX > l.Width {
		endX = l.Width
	}

	if endY > l.Height {
		endY = l.Height
	}

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			if tile := l.Get(x, y); tile != nil {
				tile.Draw(mwp)
				tileCount++
			}
		}
	}

	// for i := 0; i < len(l.Tiles); i++ {
	// 	l.Tiles[i].Draw(mwp)
	// 	tileCount++
	// }

	rl.DrawText(fmt.Sprintf("TILECOUNT: %d", tileCount), int32(camera.Target.X+5), int32(camera.Target.Y+5), 24, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("X: %d->%d, Y: %d->%d, Drawn: %d", startX, endX, startY, endY, tileCount), int32(camera.Target.X+5), int32(camera.Target.Y+30), 24, rl.DarkGray)
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
