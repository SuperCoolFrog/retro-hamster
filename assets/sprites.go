package assets

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
	Sprite_Wheel   = AssetSprite{AssetKey_Wheel_PNG, 0, 0, 1100, 1100}
	Sprite_Hamster = AssetSprite{AssetKey_Static_PNG, 0, 0, 256, 256}
)

func SpriteEquals(a AssetSprite, b AssetSprite) bool {
	return a.AssetKey == b.AssetKey &&
		a.H == b.H &&
		a.W == b.W &&
		a.X == b.W &&
		a.Y == b.Y
}
