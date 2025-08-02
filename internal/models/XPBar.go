package models

import (
	"image"
	"retro-hamster/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const NORMAL_XP_VALUE = 100.0

type XPBar struct {
	Total   float64
	Current float64

	X, Y, W, H float64

	assetStaticSpriteSheet GameAssetImg
}

func NewXP(game *Game, total, current float64) *XPBar {
	W := float64(assets.Sprite_XP_Bar_1.W)
	H := float64(assets.Sprite_XP_Bar_1.H)

	return &XPBar{
		X:                      float64(game.ScreenW)/2 - W/2,
		Y:                      float64(game.ScreenH) * 9 / 10,
		W:                      W,
		H:                      H,
		assetStaticSpriteSheet: game.ImageAssets[assets.AssetKey_Static_PNG],
		Total:                  total,
		Current:                current,
	}
}

func (m *XPBar) Update() {
	if m.Current > m.Total {
		m.Current = m.Total
	} else if m.Current < 0 {
		m.Current = 0
	}
}

func (m *XPBar) Draw(screen *ebiten.Image) {
	if m.Current > 0 {
		rat := m.Current / m.Total

		m2 := assets.Sprite_XP_Bar_2
		m2Sub := m.assetStaticSpriteSheet.Image.SubImage(m2.GetImageRect()).(*ebiten.Image)

		m2Width := int(float64(m2Sub.Bounds().Dx()) * rat)
		m2Height := m2Sub.Bounds().Dy()

		bounds := m2Sub.Bounds()
		m2Part := m2Sub.SubImage(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Min.X+m2Width, bounds.Min.Y+m2Height)).(*ebiten.Image)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(m.X, m.Y)
		screen.DrawImage(m2Part, opts)
	}

	m1 := assets.Sprite_XP_Bar_1
	m1Sub := m.assetStaticSpriteSheet.Image.SubImage(m1.GetImageRect()).(*ebiten.Image)

	m1Opts := &ebiten.DrawImageOptions{}
	m1Opts.GeoM.Translate(m.X, m.Y)

	screen.DrawImage(m1Sub, m1Opts)
}
