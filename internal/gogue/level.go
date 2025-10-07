package gogue

import (
	"fmt"
	"math/rand/v2"

	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelKind int

const (
	OVERWORLD LevelKind = iota
	CAVE
)

type Level struct {
	Tiles  []*Tile
	Width  int
	Height int
	Kind   LevelKind
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
	return t != nil && t.IsPassable
}

func (l *Level) InsertRandomDungeonDoor() {

	candidates := []*Tile{}

	for i := 0; i < len(l.Tiles); i++ {
		nbs := l.GetAllNeighbourTiles(l.Tiles[i])
		if CountWallTilesInSlice(nbs, l.Kind) == 3 {
			if l.Kind == LevelKind(OVERWORLD) {
				candidates = append(candidates, l.Tiles[i])
			}
			if l.Kind == LevelKind(CAVE) && len(nbs) == 4 {
				candidates = append(candidates, l.Tiles[i])
			}
		}
	}

	if len(candidates) == 0 {
		fmt.Println("NO VALID DOORS FOUND")
		return
	}

	idx := rand.IntN(len(candidates))
	doorTile := candidates[idx]
	doorTile.Kind = TileKind(DOOR)
	doorTile.IsPassable = true
	fmt.Printf("INSERTED DOOR: %d, %d", doorTile.Position.X, doorTile.Position.Y)
}

func GenerateCave(w, h int) *Level {
	l := Level{
		Tiles:  nil,
		Width:  w,
		Height: h,
		Kind:   LevelKind(CAVE),
	}

	alpha := 2.0
	beta := 2.0
	n := int32(3)
	seed := rand.Int64()

	p := perlin.NewPerlin(alpha, beta, n, seed)

	scale := 0.1

	for j := 0; j < l.Height; j++ {
		for i := 0; i < l.Width; i++ {

			val := p.Noise2D(float64(i)*scale, float64(j)*scale)
			// val := 0.25*p.Noise2D(float64(i)*0.05, float64(j)*0.05) + 0.75*p.Noise2D(float64(i)*0.15, float64(j)*0.15)

			val = (val + 1) / 2

			if j <= 10 || j >= l.Height-10 || i <= 10 || i >= l.Width-10 {
				l.Tiles = append(l.Tiles, &Tile{
					// Position:   MapPosition{X: (i)*TILE_SIZE + TILE_SIZE/2, Y: (j)*16 + TILE_SIZE/2},
					Position:   MapPosition{X: i, Y: j},
					Color:      rl.DarkGray,
					IsPassable: false,
					Kind:       TileKind(CAVE_WALL),
				})
			} else if val > 0.4 {
				l.Tiles = append(l.Tiles, &Tile{
					// Position:   MapPosition{X: (i)*TILE_SIZE + TILE_SIZE/2, Y: (j)*16 + TILE_SIZE/2},
					Position:   MapPosition{X: i, Y: j},
					Color:      rl.Gray,
					IsPassable: true,
					Kind:       TileKind(STONE_FLOOR),
				})
			} else {
				l.Tiles = append(l.Tiles, &Tile{
					// Position:   MapPosition{X: (i)*TILE_SIZE + TILE_SIZE/2, Y: (j)*16 + TILE_SIZE/2},
					Position:   MapPosition{X: i, Y: j},
					Color:      rl.DarkGray,
					IsPassable: false,
					Kind:       TileKind(CAVE_WALL),
				})
			}
		}
	}

	l.InsertRandomDungeonDoor()

	return &l
}

func GenerateLevel(w, h int) *Level {
	l := Level{
		Tiles:  nil,
		Width:  w,
		Height: h,
		Kind:   LevelKind(OVERWORLD),
	}

	alpha := 2.0
	beta := 2.0
	n := int32(3)
	seed := rand.Int64()

	p := perlin.NewPerlin(alpha, beta, n, seed)

	scale := 0.1

	for j := 0; j < l.Height; j++ {
		for i := 0; i < l.Width; i++ {

			val := p.Noise2D(float64(i)*scale, float64(j)*scale)
			// val := 0.25*p.Noise2D(float64(i)*0.05, float64(j)*0.05) + 0.75*p.Noise2D(float64(i)*0.15, float64(j)*0.15)

			val = (val + 1) / 2

			// if j == 10 {
			// 	l.Tiles = append(l.Tiles, &Tile{
			// 		// Position:   MapPosition{X: (i)*TILE_SIZE + TILE_SIZE/2, Y: (j)*16 + TILE_SIZE/2},
			// 		Position:   MapPosition{X: i, Y: j},
			// 		Color:      rl.Gray,
			// 		IsPassable: false,
			// 		Kind:       TileKind(WALL),
			// 	})
			// } else
			if val > 0.4 {
				l.Tiles = append(l.Tiles, &Tile{
					// Position:   MapPosition{X: (i)*TILE_SIZE + TILE_SIZE/2, Y: (j)*16 + TILE_SIZE/2},
					Position:   MapPosition{X: i, Y: j},
					Color:      rl.Green,
					IsPassable: true,
					Kind:       TileKind(GRASS),
				})
			} else {
				l.Tiles = append(l.Tiles, &Tile{
					// Position:   MapPosition{X: (i)*TILE_SIZE + TILE_SIZE/2, Y: (j)*16 + TILE_SIZE/2},
					Position:   MapPosition{X: i, Y: j},
					Color:      rl.Gray,
					IsPassable: false,
					Kind:       TileKind(WALL),
				})
			}
		}
	}

	l.InsertRandomDungeonDoor()

	return &l
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
		nX := t.Position.X + d.X
		nY := t.Position.Y + d.Y
		n := l.Get(nX, nY)

		if n != nil {
			//n.Color = rl.Orange
			neighbs = append(neighbs, *n)
		}

	}

	//fmt.Printf("len(neighbs): %d\n", len(neighbs))
	return neighbs
}

func CountWallTilesInSlice(s []Tile, lk LevelKind) int {
	var count = 0
	for _, t := range s {
		// fmt.Println(t.Kind)

		if lk == LevelKind(OVERWORLD) && t.Kind == TileKind(WALL) {
			count += 1
			// fmt.Printf("WALL NBS: %d\n", count)
		}

		if lk == LevelKind(CAVE) && t.Kind == TileKind(CAVE_WALL) {
			count += 1
			// fmt.Printf("WALL NBS: %d\n", count)
		}
	}

	return count
}
