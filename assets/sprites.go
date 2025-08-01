package assets

import "image"

type AssetSize struct {
	W int
	H int
}

type AssetSprite struct {
	AssetKey AssetKey
	X        int
	Y        int
	W        int
	H        int
}

var (
	Sprite_Wheel          = AssetSprite{AssetKey_Wheel_PNG, 0, 0, 1100, 1100}
	Sprite_Hamster        = AssetSprite{AssetKey_Static_PNG, 0, 0, 256, 256}
	Sprite_Momentum_Bar_1 = AssetSprite{AssetKey_Static_PNG, 256, 0, 256, 64}
	Sprite_Momentum_Bar_2 = AssetSprite{AssetKey_Static_PNG, 256, 64, 256, 64}
)

func (spr AssetSprite) GetImageRect() image.Rectangle {
	return image.Rect(spr.X, spr.Y, spr.X+spr.W, spr.Y+spr.H)
}

func SpriteEquals(a AssetSprite, b AssetSprite) bool {
	return a.AssetKey == b.AssetKey &&
		a.H == b.H &&
		a.W == b.W &&
		a.X == b.W &&
		a.Y == b.Y
}
