package assets

import (
	_ "embed"
)

type AssetKey = string
type AssetType = uint8

const (
	AssetType_PNG AssetType = 1 << 0
	AssetType_TTF AssetType = 1 << 1
)

type AssetRef struct {
	Key       AssetKey
	AssetType AssetType
	Data      []byte
}

var Assets = map[AssetKey]AssetRef{}
