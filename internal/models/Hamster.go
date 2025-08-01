package models

import (
	"retro-hamster/assets"
	"retro-hamster/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hamster struct {
	X, Y         float64
	W, H         float64
	IsRunning    bool
	Direction    DIRECTION
	AnimationRun *Animation

	assetStaticSpriteSheet GameAssetImg
	assetRunSpriteSheet    GameAssetImg
}

func NewHamster(game *Game) *Hamster {
	X := float64(game.ScreenW/2 - assets.AnimationHamsterRun.InitialSprite.W/2)
	Y := float64(game.ScreenH/2 - int(float64(assets.AnimationHamsterRun.InitialSprite.H)/1.25))
	return &Hamster{
		X:                      X,
		Y:                      Y,
		W:                      float64(assets.AnimationHamsterRun.InitialSprite.W),
		H:                      float64(assets.AnimationHamsterRun.InitialSprite.H),
		assetStaticSpriteSheet: game.ImageAssets[assets.AssetKey_Static_PNG],
		assetRunSpriteSheet:    game.ImageAssets[assets.AssetKey_Hamster_Run_PNG],
		AnimationRun: &Animation{
			FPS:          12,
			CurrentFrame: 0,
			Details:      assets.AnimationHamsterRun,
			X:            X,
			Y:            Y,
		},
		IsRunning: false,
		Direction: DIRECTION_RIGHT,
	}
}

func (s *Hamster) Update() {
	if !s.IsRunning {
		s.AnimationRun.CurrentFrame = 0
	} else {
		s.AnimationRun.AdvanceFrame()
	}
}

func (s *Hamster) Draw(screen *ebiten.Image) {
	// DrawCollisionRect(screen, s.GetCollisionRect(), color.RGBA{0, 255, 0, 255})

	if !s.IsRunning {
		s.drawStaticHamster(screen)
		return
	}

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

func (s *Hamster) GetCollisionRect() CollisionRect {
	return CollisionRect{s.X, s.Y, s.W, s.H}
}
