package levels

const LEVEL_1_CHART = `------------------------
_o_o_o_o_M_o_o_o
_o_S_o_m_o_Q_o_o`

const LEVEL_2_CHART = `------------------------
_o_M_o_o_S_o_o_o_o_S
_o o o o SS o S o o o`

const LEVEL_3_CHART = `------------------------
_o_o_o_o_Q_m_Q_o_S_o_
_Q o o o QQ o S o o o`

const LEVEL_4_CHART = `------------------------
_o_o_|_o_S_m_Q_o_S_o_
_Q o | o M o S o o o`

var ALL_LEVEL_CHARTS = []string{
	LEVEL_1_CHART,
	LEVEL_2_CHART,
	LEVEL_3_CHART,
	LEVEL_4_CHART,
}

const LEVEL_BOSS_1 = `------------------------
_SS_Q_o_o_o_o_o_o_
_o_o_o_o_B_o_o_S_o`

const LEVEL_BOSS_2 = `------------------------
_m___m_o_2_o_o_m_o`

const LEVEL_BOSS_3 = `------------------------
_______|3|______`

var BOSS_LEVELS = []string{
	LEVEL_BOSS_1,
	LEVEL_BOSS_2,
	LEVEL_BOSS_3,
}
