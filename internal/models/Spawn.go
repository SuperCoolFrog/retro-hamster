package models

import (
	"math"
)

type Spawn struct {
	SpawnAngle     float64 // initial angle on the ground (radians)
	WheelRadius    float64 // distance from center
	SpawnAnimation Animation

	// X, Y float64

	spawnRotation float64
}

func NewSpawn(spawnAngle float64, wheelRadius float64, spawnAnimation *Animation) *Spawn {
	return &Spawn{
		SpawnAngle:     spawnAngle,
		WheelRadius:    wheelRadius,
		SpawnAnimation: *spawnAnimation,
	}
}

func (s *Spawn) Update(wheelCenterX, wheelCenterY, wheelAngle float64) {
	angle := s.SpawnAngle + wheelAngle
	s.SpawnAnimation.X = wheelCenterX + s.WheelRadius*math.Cos(angle)
	s.SpawnAnimation.Y = wheelCenterY + s.WheelRadius*math.Sin(angle)

	s.spawnRotation = s.SpawnAngle + wheelAngle + math.Pi/2

	s.SpawnAnimation.AdvanceFrame()
}

// func (obj *Spawn) Draw(screen *ebiten.Image, img *ebiten.Image, cx, cy, groundAngle float64) {
// 	angle := obj.spawnAngle + groundAngle
// 	x := cx + obj.radius*math.Cos(angle)
// 	y := cy + obj.radius*math.Sin(angle)

// 	rotation := angle + math.Pi/2 // Standing upright on the ground

// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(-float64(img.Bounds().Dx())/2, -float64(img.Bounds().Dy())/2)
// 	op.GeoM.Rotate(rotation)
// 	op.GeoM.Translate(x, y)

// 	screen.DrawImage(img, op)
// }
