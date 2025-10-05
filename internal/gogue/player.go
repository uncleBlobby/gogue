package gogue

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Position   rl.Vector2
	Speed      float32
	MoveTarget Tile
}

func (p *Player) IsAtMoveTarget() bool {
	if p.Position.X != p.MoveTarget.Position.X || p.Position.Y != p.MoveTarget.Position.Y {
		return false
	}

	return true
}

func (p *Player) Update(dt float32, l Level, mwp rl.Vector2) {

	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		for _, t := range l.Tiles {
			if rl.CheckCollisionPointRec(mwp, rl.Rectangle{
				X:      t.Position.X,
				Y:      t.Position.Y,
				Width:  TILE_SIZE,
				Height: TILE_SIZE,
			}) {
				p.MoveTarget = t
			}
		}
	}

	if !p.IsAtMoveTarget() {

		moveDir := rl.Vector2Subtract(p.MoveTarget.Position, p.Position)
		// move toward target

		p.Position.X += moveDir.X * p.Speed * dt
		p.Position.Y += moveDir.Y * p.Speed * dt
		// find path
	}
}

func (p *Player) Draw() {
	rl.DrawRectangle(int32(p.Position.X), int32(p.Position.Y), TILE_SIZE, TILE_SIZE, rl.Blue)
}
