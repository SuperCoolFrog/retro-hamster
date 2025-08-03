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
	Sprite_Background     = AssetSprite{AssetKey_Background_PNG, 0, 0, 1920, 1080}
	Sprite_Start          = AssetSprite{AssetKey_Start_PNG, 0, 0, 1920, 1080}
	Sprite_Wheel          = AssetSprite{AssetKey_Wheel_PNG, 0, 0, 1100, 1100}
	Sprite_Hamster        = AssetSprite{AssetKey_Static_PNG, 0, 0, 256, 256}
	Sprite_Momentum_Bar_1 = AssetSprite{AssetKey_Static_PNG, 256, 0, 256, 64}
	Sprite_Momentum_Bar_2 = AssetSprite{AssetKey_Static_PNG, 256, 64, 256, 64}
	Sprite_Heart          = AssetSprite{AssetKey_Static_PNG, 0, 960, 64, 64}
	Sprite_Seed           = AssetSprite{AssetKey_Static_PNG, 128, 960, 64, 64}
	Sprite_XP_Bar_1       = AssetSprite{AssetKey_Static_PNG, 0, 864, 768, 64}
	Sprite_XP_Bar_2       = AssetSprite{AssetKey_Static_PNG, 0, 800, 768, 64}
	Sprite_Block          = AssetSprite{AssetKey_Static_PNG, 512, 0, 128, 128}
	Sprite_Fence          = AssetSprite{AssetKey_Static_PNG, 640, 0, 128, 256}
	Sprite_Button         = AssetSprite{AssetKey_Static_PNG, 0, 384, 256, 128}
	Sprite_Button_Focused = AssetSprite{AssetKey_Static_PNG, 0, 512, 256, 128}
	Sprite_Health_Card    = AssetSprite{AssetKey_Static_PNG, 384, 384, 256, 256}
	Sprite_Momentum_Card  = AssetSprite{AssetKey_Static_PNG, 640, 384, 256, 256}
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
