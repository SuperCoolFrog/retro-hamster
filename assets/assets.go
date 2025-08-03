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
	AssetKey_NONE            AssetKey = "AssetKey_NONE"
	AssetKey_Static_PNG      AssetKey = "AssetKey_Static_PNG"
	AssetKey_Wheel_PNG       AssetKey = "AssetKey_Wheel_PNG"
	AssetKey_Hamster_Run_PNG AssetKey = "AssetKey_Hamster_Run_PNG"
	AssetKey_Snake_PNG       AssetKey = "AssetKey_Snake_PNG"
	AssetKey_Background_PNG  AssetKey = "AssetKey_Background_PNG"
	AssetKey_Start_PNG       AssetKey = "AssetKey_Start_PNG"
	AssetKey_Shark_PNG       AssetKey = "AssetKey_Shark_PNG"
)

var (
	//go:embed wheel.png
	m_wheel_png []byte

	//go:embed hamster_run.png
	m_hamster_run_png []byte

	//go:embed static.png
	m_static_png []byte

	//go:embed snake.png
	m_snake_png []byte

	//go:embed shark.png
	m_shark_png []byte

	//go:embed background.png
	m_background_png []byte

	//go:embed start.png
	m_start_png []byte
)

var Assets = map[AssetKey]AssetRef{
	AssetKey_Wheel_PNG: {
		Key:       AssetKey_Wheel_PNG,
		AssetType: AssetType_PNG,
		Data:      m_wheel_png,
	},
	AssetKey_Hamster_Run_PNG: {
		Key:       AssetKey_Hamster_Run_PNG,
		AssetType: AssetType_PNG,
		Data:      m_hamster_run_png,
	},
	AssetKey_Static_PNG: {
		Key:       AssetKey_Static_PNG,
		AssetType: AssetType_PNG,
		Data:      m_static_png,
	},
	AssetKey_Snake_PNG: {
		Key:       AssetKey_Snake_PNG,
		AssetType: AssetType_PNG,
		Data:      m_snake_png,
	},
	AssetKey_Background_PNG: {
		Key:       AssetKey_Background_PNG,
		AssetType: AssetType_PNG,
		Data:      m_background_png,
	},
	AssetKey_Start_PNG: {
		Key:       AssetKey_Start_PNG,
		AssetType: AssetType_PNG,
		Data:      m_start_png,
	},
	AssetKey_Shark_PNG: {
		Key:       AssetKey_Shark_PNG,
		AssetType: AssetType_PNG,
		Data:      m_shark_png,
	},
}
