package models

import (
	"retro-hamster/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameAssetImg struct {
	Asset assets.AssetRef
	Image *ebiten.Image
}

type GameAssetFont struct {
	Asset assets.AssetRef
	Font  *text.GoTextFaceSource
}
