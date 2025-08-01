package models

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CollisionRect struct {
	X, Y float64
	W, H float64
}

func (r CollisionRect) Intersects(other CollisionRect) bool {
	return r.X < other.X+other.W &&
		r.X+r.W > other.X &&
		r.Y < other.Y+other.H &&
		r.Y+r.H > other.Y
}

func DrawCollisionRect(screen *ebiten.Image, r CollisionRect, col color.Color) {
	screen.Set(int(r.X), int(r.Y), col)
	ebitenutil.DrawRect(screen, r.X, r.Y, r.W, r.H, col)
}
