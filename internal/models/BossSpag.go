package models

type BossSpag struct {
	HasInit bool
	Health  float64
}

var BOSS_SPAG = &BossSpag{
	HasInit: false,
	Health:  10,
}
