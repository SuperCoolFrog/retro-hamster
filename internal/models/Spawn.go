package models

import (
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Spawn struct {
	SpawnAngle     float64 // initial angle on the ground (radians)
	WheelRadius    float64 // distance from center
	SpawnAnimation Animation
	IsAlive        bool
	IsObstacle     bool
	Direction      DIRECTION
	OnCollision    func(ham *Hamster)

	LastActivation     time.Time
	ActivationCoolDown time.Duration

	X, Y float64
	W, H float64

	Power float64

	spawnRotation float64
}

func NewSpawn(spawnAngle float64, wheelRadius float64, spawnAnimation *Animation) *Spawn {
	return &Spawn{
		SpawnAngle:     spawnAngle,
		WheelRadius:    wheelRadius,
		SpawnAnimation: *spawnAnimation,
		Direction:      DIRECTION_RIGHT,
		W:              float64(spawnAnimation.Details.InitialSprite.W),
		H:              float64(spawnAnimation.Details.InitialSprite.H),
		IsAlive:        true,
		OnCollision:    func(ham *Hamster) {},
	}
}

func (s *Spawn) SetHamsterRelativeOffset(angle float64) {
	s.SpawnAngle = s.SpawnAngle - angle
}

func (s *Spawn) Update(wheelCenterX, wheelCenterY, wheelAngle float64) {
	angle := s.SpawnAngle + wheelAngle

	s.X = wheelCenterX + s.WheelRadius*math.Cos(angle)
	s.Y = wheelCenterY + s.WheelRadius*math.Sin(angle)

	s.SpawnAnimation.X = s.X
	s.SpawnAnimation.Y = s.Y

	s.spawnRotation = s.SpawnAngle + wheelAngle + math.Pi/2

	s.SpawnAnimation.AdvanceFrame()
}

func (s *Spawn) Draw(game *Game, screen *ebiten.Image) {
	// s.drawHitBox(screen)

	img := s.SpawnAnimation.GetCurrentFrame()
	animSs := game.ImageAssets[img.AssetKey]
	asset := img.AssetSprite

	sub := animSs.Image.SubImage(image.Rect(asset.X, asset.Y, asset.X+asset.W, asset.Y+asset.H)).(*ebiten.Image)

	hOpts := ebiten.DrawImageOptions{}
	if s.Direction == DIRECTION_LEFT {
		hOpts.GeoM.Scale(float64(s.Direction), 1)
		hOpts.GeoM.Translate(float64(sub.Bounds().Dx()), 0) // Corrects after flipping
	}

	hOpts.GeoM.Translate(-float64(sub.Bounds().Dx())/2, -float64(sub.Bounds().Dy())/2)
	// hOpts.GeoM.Translate(-float64(s.W)/2, -float64(s.H)/2)
	hOpts.GeoM.Rotate(s.spawnRotation)

	hOpts.GeoM.Translate(float64(img.TargetX), float64(img.TargetY))
	hOpts.Filter = ebiten.FilterLinear

	hOpts.GeoM.Translate(float64(sub.Bounds().Dx()), 0) // Corrects after flipping
	screen.DrawImage(sub,
		&hOpts,
	)
}

func (s *Spawn) GetHitBox() [4]Vector {
	hw := s.W / 2
	hh := s.H / 2

	// Corner positions relative to center
	corners := [4]Vector{
		{-hw, -hh}, // top-left
		{+hw, -hh}, // top-right
		{+hw, +hh}, // bottom-right
		{-hw, +hh}, // bottom-left
	}

	sin, cos := math.Sin(s.spawnRotation), math.Cos(s.spawnRotation)
	for i := range corners {
		cx := corners[i].X + s.W
		cy := corners[i].Y
		corners[i].X = s.X + cx*cos - cy*sin
		corners[i].Y = s.Y + cx*sin + cy*cos
	}

	return corners
}
