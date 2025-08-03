package models

import (
	"image"
	"math"
	"retro-hamster/assets"
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
	ModHitBox      float64
	Health         float64
	SkewAngle      float64

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
		ModHitBox:      1,
		SkewAngle:      -1,
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
	DrawHitBox(screen, s.GetHitBox())
	DrawHitBox(screen, s.GetRenderQuad())

	img := s.SpawnAnimation.GetCurrentFrame()
	animSs := game.ImageAssets[img.AssetKey]
	asset := img.AssetSprite

	opts := QuadDrawOptions{
		FlipX:   s.Direction != DIRECTION(s.SpawnAnimation.Details.Direction) && s.SpawnAnimation.Details.Direction != assets.ANIMATION_DIRECTION_NONE,
		FlipY:   false,
		SrcRect: asset.GetImageRect(),
	}

	s.DrawImageFromQuad(screen, animSs.Image, opts)
}

func (s *Spawn) DrawImageFromQuad(screen *ebiten.Image, spriteSheet *ebiten.Image, opts QuadDrawOptions) {
	hitbox := s.GetRenderQuad()

	src := opts.SrcRect
	texW := float32(src.Dx())
	texH := float32(src.Dy())

	// Handle flipping of UVs
	var u0, u1 float32 = 0, texW
	var v0, v1 float32 = 0, texH
	if opts.FlipX {
		u0, u1 = texW, 0
	}
	if opts.FlipY {
		v0, v1 = texH, 0
	}

	vertices := []ebiten.Vertex{
		{
			DstX: float32(hitbox[0].X),
			DstY: float32(hitbox[0].Y),
			SrcX: float32(src.Min.X) + u0, SrcY: float32(src.Min.Y) + v0,
			ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1,
		},
		{
			DstX: float32(hitbox[1].X),
			DstY: float32(hitbox[1].Y),
			SrcX: float32(src.Min.X) + u1, SrcY: float32(src.Min.Y) + v0,
			ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1,
		},
		{
			DstX: float32(hitbox[2].X),
			DstY: float32(hitbox[2].Y),
			SrcX: float32(src.Min.X) + u1, SrcY: float32(src.Min.Y) + v1,
			ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1,
		},
		{
			DstX: float32(hitbox[3].X),
			DstY: float32(hitbox[3].Y),
			SrcX: float32(src.Min.X) + u0, SrcY: float32(src.Min.Y) + v1,
			ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1,
		},
	}

	indices := []uint16{
		0, 1, 2,
		0, 2, 3,
	}

	triangleOpts := &ebiten.DrawTrianglesOptions{
		Filter: ebiten.FilterLinear,
	}

	screen.DrawTriangles(vertices, indices, spriteSheet, triangleOpts)
}

func (s *Spawn) GetHitBox() [4]Vector {
	hw := (s.W) / 2
	hh := (s.H) / 2

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

	scaled := s.scaleHitbox(corners)
	if s.SkewAngle != -1 {
		return RotateQuad(scaled, s.SkewAngle)
	}

	return scaled
}

func (s *Spawn) GetRenderQuad() [4]Vector {
	hw := (s.W) / 2
	hh := (s.H) / 2

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

	if s.SkewAngle != -1 {
		return RotateQuad(corners, s.SkewAngle)
	}

	return corners
}

type QuadDrawOptions struct {
	FlipX   bool            // Flip horizontally
	FlipY   bool            // Flip vertically
	SrcRect image.Rectangle // Subsection of sprite sheet to draw (sprite frame)
}

func (s *Spawn) scaleHitbox(hitbox [4]Vector) [4]Vector {
	scaleX := s.ModHitBox
	scaleY := s.ModHitBox
	// // Step 1: Compute center of the quad
	// var cx, cy float64
	// for _, v := range hitbox {
	// 	cx += v.X
	// 	cy += v.Y
	// }
	// cx /= 4
	// cy /= 4

	// // Step 2: Scale each point relative to the center
	// var scaled [4]Vector
	// for i, v := range hitbox {
	// 	dx := v.X - cx
	// 	dy := v.Y - cy
	// 	scaled[i] = Vector{
	// 		X: cx + dx*scaleX,
	// 		Y: cy + dy*scaleY,
	// 	}
	// }

	// return scaled

	// Step 1: Compute bottom-center point
	bottomMid := Vector{
		X: (hitbox[2].X + hitbox[3].X) / 2,
		Y: (hitbox[2].Y + hitbox[3].Y) / 2,
	}

	// Step 2: Scale each point relative to bottom-center
	var scaled [4]Vector
	for i, v := range hitbox {
		dx := v.X - bottomMid.X
		dy := v.Y - bottomMid.Y
		scaled[i] = Vector{
			X: bottomMid.X + dx*scaleX,
			Y: bottomMid.Y + dy*scaleY,
		}
	}

	return scaled
}

func RotateQuad(quad [4]Vector, angle float64) [4]Vector {
	// Step 1: Compute center of quad
	var cx, cy float64
	for _, v := range quad {
		cx += v.X
		cy += v.Y
	}
	cx /= 4
	cy /= 4

	// Step 2: Rotate each point around the center
	sinA, cosA := math.Sin(angle), math.Cos(angle)
	var rotated [4]Vector
	for i, v := range quad {
		dx := v.X - cx
		dy := v.Y - cy
		rotated[i] = Vector{
			X: cx + dx*cosA - dy*sinA,
			Y: cy + dx*sinA + dy*cosA,
		}
	}

	return rotated
}
