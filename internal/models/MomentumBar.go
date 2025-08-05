package models

import (
	"image"
	"retro-hamster/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const NORMAL_MOMENTUM_VALUE = 100.0

type MomentumBar struct {
	Total   float64
	Current float64

	X, Y, W, H float64

	assetStaticSpriteSheet GameAssetImg
}

func NewMomentumBar(game *Game, total, current float64) *MomentumBar {
	W := float64(assets.Sprite_Momentum_Bar_1.W)
	H := float64(assets.Sprite_Momentum_Bar_1.H)

	return &MomentumBar{
		X:                      float64(game.ScreenW) / 2, // - W/2,
		Y:                      float64(game.ScreenH) * 3 / 4,
		W:                      W,
		H:                      H,
		assetStaticSpriteSheet: game.ImageAssets[assets.AssetKey_Static_PNG],
		Total:                  total,
		Current:                current,
	}
}

func (m *MomentumBar) Update() {
	if m.Current > m.Total {
		m.Current = m.Total
	} else if m.Current < 0 {
		m.Current = 0
	}
}

func (m *MomentumBar) Draw(screen *ebiten.Image) {
	scale := m.Total / NORMAL_MOMENTUM_VALUE

	if m.Current > 0 {
		rat := m.Current / m.Total

		m2 := assets.Sprite_Momentum_Bar_2
		m2Sub := m.assetStaticSpriteSheet.Image.SubImage(m2.GetImageRect()).(*ebiten.Image)

		m2Width := int(float64(m2Sub.Bounds().Dx()) * rat)
		m2Height := m2Sub.Bounds().Dy()

		bounds := m2Sub.Bounds()
		m2Part := m2Sub.SubImage(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Min.X+m2Width, bounds.Min.Y+m2Height)).(*ebiten.Image)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(scale, 1)
		opts.GeoM.Translate(m.X-(float64(m2.W)*scale)/2, m.Y-(float64(m2.H)/2))
		screen.DrawImage(m2Part, opts)
	}

	m1 := assets.Sprite_Momentum_Bar_1
	m1Sub := m.assetStaticSpriteSheet.Image.SubImage(m1.GetImageRect()).(*ebiten.Image)
	// Calculate the image's original center
	m1Width := m1Sub.Bounds().Dx()
	m1Height := m1Sub.Bounds().Dy()

	m1Opts := &ebiten.DrawImageOptions{}

	// Translate to move the center to the origin (0,0)
	m1Opts.GeoM.Translate(float64(-m1Width)/2, float64(-m1Height)/2)

	m1Opts.GeoM.Scale(scale, 1)
	m1Opts.GeoM.Translate(m.X, m.Y)

	screen.DrawImage(m1Sub, m1Opts)
}

func (m *MomentumBar) Ratio() float64 {
	return m.Current / m.Total
}
