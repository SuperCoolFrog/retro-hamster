package levels

const LEVEL_1_CHART = `------------------------
_o_o_o_o_M_o_o_o
_o_S_o_m_o_Q_o_o`

const LEVEL_2_CHART = `------------------------
_o_o_o_o_S_o_o_o_o_S
_o o o o SS o S o o o`

var ALL_LEVEL_CHARTS = []string{
	LEVEL_1_CHART,
	// LEVEL_2_CHART,
}

const LEVEL_BOSS_1 = `------------------------
_SS_Q_o_o_o_o_o_o_
_o_o_o_o_B_o_o_S_o`

const LEVEL_BOSS_2 = `------------------------
_m___m_o_2_o_o_m_o`

const LEVEL_BOSS_3 = `------------------------
_______|_3|______`

var BOSS_LEVELS = []string{
	LEVEL_BOSS_1,
	LEVEL_BOSS_2,
	LEVEL_BOSS_3,
}
