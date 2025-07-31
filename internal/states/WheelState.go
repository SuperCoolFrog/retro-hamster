package states

import (
	"math"
	"retro-hamster/assets"
	"retro-hamster/internal/models"
	"retro-hamster/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type WheelState struct {
	Game  *models.Game
	angle float64
}

func (s *WheelState) OnTransition() {}

func (s *WheelState) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.angle -= 2 * math.Pi / 180
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.angle += 2 * math.Pi / 180
	}
	return nil
}

func (s *WheelState) Draw(screen *ebiten.Image) {
	wheelPng := s.Game.ImageAssets[assets.AssetKey_Wheel_PNG]
	op := ebiten.DrawImageOptions{}

	scale := 2.0
	wheelW := float64(assets.Sprite_Wheel.W) * scale
	wheelH := float64(assets.Sprite_Wheel.H) * scale

	// Move the origin to the center of the image before rotating
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(-wheelW/2, -wheelH/2)
	op.GeoM.Rotate(s.angle)
	op.Filter = ebiten.FilterLinear
	// Move it to the screen center after rotation
	op.GeoM.Translate(float64(s.Game.ScreenW)/2, float64(s.Game.ScreenH)/2+wheelH/2)

	scenes.DrawAssetSpriteWithOptions(wheelPng.Image, screen, assets.Sprite_Wheel, op)
}
