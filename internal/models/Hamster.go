package models

import (
	"math"
	"retro-hamster/assets"
	"retro-hamster/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hamster struct {
	X, Y          float64
	W, H          float64
	IsRunning     bool
	IsJumping     bool
	OriginalAngle float64
	LogicalAngle  float64

	Direction    DIRECTION
	AnimationRun *Animation

	assetStaticSpriteSheet GameAssetImg
	assetRunSpriteSheet    GameAssetImg

	Momentum *MomentumBar
	XP       *XPBar
	Health   int

	gravity       float64
	vY            float64
	initialY      float64
	lastDirection DIRECTION
}

func NewHamster(game *Game) *Hamster {
	X := float64(game.ScreenW/2 - assets.AnimationHamsterRun.InitialSprite.W/2)
	Y := float64(game.ScreenH/2 - int(float64(assets.AnimationHamsterRun.InitialSprite.H)/1.25))
	return &Hamster{
		X:                      X,
		Y:                      Y,
		W:                      float64(assets.AnimationHamsterRun.InitialSprite.W),
		H:                      float64(assets.AnimationHamsterRun.InitialSprite.H),
		initialY:               Y,
		OriginalAngle:          math.Pi / 2,
		LogicalAngle:           math.Pi / 2,
		assetStaticSpriteSheet: game.ImageAssets[assets.AssetKey_Static_PNG],
		assetRunSpriteSheet:    game.ImageAssets[assets.AssetKey_Hamster_Run_PNG],
		AnimationRun: &Animation{
			FPS:          12,
			CurrentFrame: 0,
			Details:      assets.AnimationHamsterRun,
			X:            X,
			Y:            Y,
		},
		IsRunning:     false,
		Direction:     DIRECTION_RIGHT,
		lastDirection: DIRECTION_RIGHT,
		gravity:       1,
		Momentum:      NewMomentumBar(game, 100, 0),
		XP:            NewXP(game, 100, 0),
		Health:        3,
	}
}

func (s *Hamster) InitJump() {
	if !s.IsJumping {
		s.IsJumping = true
		s.vY = -20
	}
}

func (s *Hamster) Update(wheelAngle float64) {
	s.LogicalAngle = s.OriginalAngle + wheelAngle

	if s.lastDirection != s.Direction {
		s.Momentum.Current = 0
		s.lastDirection = s.Direction
	}

	if !s.IsRunning {
		s.AnimationRun.CurrentFrame = 0
		s.Momentum.Current -= 3
	} else {
		s.AnimationRun.AdvanceFrame()
		s.Momentum.Current += 1
	}

	if s.IsJumping {
		s.Y += s.vY
		s.vY += s.gravity

		if s.Y >= s.initialY {
			s.IsJumping = false
			s.Y = s.initialY
			s.vY = 0
		}
	}

	s.Momentum.Update()
	s.XP.Update()
}

func (s *Hamster) Draw(screen *ebiten.Image) {
	// DrawCollisionRect(screen, s.GetCollisionRect(), color.RGBA{0, 255, 0, 255})

	if !s.IsRunning || s.IsJumping {
		s.drawStaticHamster(screen)
	} else if s.IsRunning {
		runFrame := s.AnimationRun.GetCurrentFrame()
		if s.Direction == DIRECTION_LEFT {
			hOpts := ebiten.DrawImageOptions{}
			hOpts.GeoM.Scale(float64(s.Direction), 1)
			hOpts.GeoM.Translate(float64(s.X), float64(s.Y))
			scenes.DrawAssetSpriteWithOptionsWithBoundsCorrect(s.assetRunSpriteSheet.Image, screen, runFrame.AssetSprite, hOpts)
		} else {
			scenes.DrawSprite(s.assetRunSpriteSheet.Image, screen, runFrame.TargetX, runFrame.TargetY, runFrame.X, runFrame.Y, runFrame.W, runFrame.H)
		}
	}

	s.Momentum.Draw(screen)
	s.XP.Draw(screen)

	s.drawHealth(screen)
}

func (s *Hamster) drawStaticHamster(screen *ebiten.Image) {
	if s.Direction == DIRECTION_RIGHT {
		scenes.DrawAssetSprite(s.assetStaticSpriteSheet.Image, screen, s.X, s.Y, assets.Sprite_Hamster)
	} else {
		hOpts := ebiten.DrawImageOptions{}
		hOpts.GeoM.Scale(float64(s.Direction), 1)
		hOpts.GeoM.Translate(float64(s.X), float64(s.Y))
		scenes.DrawAssetSpriteWithOptionsWithBoundsCorrect(s.assetStaticSpriteSheet.Image, screen, assets.Sprite_Hamster, hOpts)
	}
}

func (s *Hamster) drawHealth(screen *ebiten.Image) {

	totalW := assets.Sprite_Heart.W * s.Health

	y := float64(s.initialY) + s.H*1.35
	startX := s.X + s.W/2 - float64(totalW)/2

	for i := range s.Health {
		x := startX + float64(i*assets.Sprite_Heart.W)
		scenes.DrawAssetSprite(s.assetStaticSpriteSheet.Image, screen, x, y, assets.Sprite_Heart)
	}
}

func (s *Hamster) GetCollisionRect() CollisionRect {
	if s.IsJumping {

		return CollisionRect{s.X, s.Y, s.W, s.H / 3}
	}
	return CollisionRect{s.X, s.Y, s.W, s.H}
}
