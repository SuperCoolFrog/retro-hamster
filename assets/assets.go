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

const (
	AssetKey_NONE      AssetKey = "AssetKey_NONE"
	AssetKey_Wheel_PNG AssetKey = "AssetKey_Wheel_PNG"
)

var (
	//go:embed wheel.png
	m_wheel_png []byte
)

var Assets = map[AssetKey]AssetRef{
	AssetKey_Wheel_PNG: {
		Key:       AssetKey_Wheel_PNG,
		AssetType: AssetType_PNG,
		Data:      m_wheel_png,
	},
}
