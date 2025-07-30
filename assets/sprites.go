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

func SpriteEquals(a AssetSprite, b AssetSprite) bool {
	return a.AssetKey == b.AssetKey &&
		a.H == b.H &&
		a.W == b.W &&
		a.X == b.W &&
		a.Y == b.Y
}
