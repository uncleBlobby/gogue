package gogue

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	Tiles  []*Tile
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
}

func (l *Level) Index(x, y int) int {
	return y*l.Width + x
}

func (l *Level) Get(x, y int) *Tile {
	if x < 0 || y < 0 || x >= l.Width || y >= l.Height {
		return nil
	}

	return l.Tiles[l.Index(x, y)]
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

func (l *Level) InsertRandomDungeonDoor() {

	// find a random wall tile somewhere that was walls on 3 sides and grass on another

	// for _, t := range l.Tiles {
	for i := 0; i < len(l.Tiles); i++ {

		nbs := l.GetAllNeighbourTiles(l.Tiles[i])

		if CountWallTilesInSlice(nbs) == 3 {
			//make door tile type
			// fmt.Printf("SPAWNING DOOR TILE X:%d Y:%d\n", l.Tiles[i].Position.X, l.Tiles[i].Position.Y)
			l.Tiles[i].Kind = TileKind(DOOR)
			//return
		}
	}

}

func (l *Level) GetAllNeighbourTiles(t *Tile) []Tile {

	var directions = []MapPosition{
		{0, 1}, // down

		{1, 0}, // right

		{0, -1}, // up

		{-1, 0}, // left

		// {1, 1}, // down right

		// {1, -1}, // up right

		// {-1, -1}, // up left

		// {-1, 1}, // down left
	}

	neighbs := make([]Tile, 0, len(directions))

	for _, d := range directions {
		nX := (t.Position.X / TILE_SIZE) + d.X
		nY := (t.Position.Y / TILE_SIZE) + d.Y
		n := l.Get(nX, nY)

		if n != nil {
			//n.Color = rl.Orange
			neighbs = append(neighbs, *n)
		}

	}

	//fmt.Printf("len(neighbs): %d\n", len(neighbs))
	return neighbs
}

func CountWallTilesInSlice(s []Tile) int {
	var count = 0
	for _, t := range s {
		// fmt.Println(t.Kind)
		if t.Kind == TileKind(WALL) {
			count += 1
			// fmt.Printf("WALL NBS: %d\n", count)
		}
	}

	return count
}
