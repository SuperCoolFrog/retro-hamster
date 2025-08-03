package models

import (
	"image/color"
	"retro-hamster/assets"
)

type DIRECTION = int8

const DIRECTION_NONE DIRECTION = 0
const DIRECTION_LEFT DIRECTION = -1
const DIRECTION_RIGHT DIRECTION = 1

const WHEEL_SCALE = 2.0
const WHEEL_RADIUS = 1100.0

var WHEEL_WIDTH = float64(assets.Sprite_Wheel.W) * WHEEL_SCALE
var WHEEL_HEIGHT = float64(assets.Sprite_Wheel.H) * WHEEL_SCALE

const SPAWN_SPACING = 256.0

var COLOR_PINK = color.RGBA{252, 48, 125, 255}
