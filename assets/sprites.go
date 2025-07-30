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
	Sprite_Wheel = AssetSprite{AssetKey_Wheel_PNG, 0, 0, 1080, 1080}
)

func SpriteEquals(a AssetSprite, b AssetSprite) bool {
	return a.AssetKey == b.AssetKey &&
		a.H == b.H &&
		a.W == b.W &&
		a.X == b.W &&
		a.Y == b.Y
}
