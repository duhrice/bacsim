package stat

import (

)

type (
	Stat int
	Stats [StatsLen]float64	// Array, not slice (since we have a fixed size)
)

const (
	NoStat Stat = iota
	BaseHp				// Base HP
	FlatHp				// Flat HP
	HpPer				// HP percent
	BaseAtk				// Base attack
	FlatAtk				// Flat attack
	AtkPer				// Attack percent
	BaseDef				// Base defense
	FlatDef				// Flat defense
	DefPer				// Defense percent
	BaseHeal			// Base healing
	FlatHeal			// Flat healing
	HealPer				// Healing percent
	BaseAcc				// Base accuracy
	FlatAcc				// Flat accuracy
	AccPer				// Accuracy percent
	BaseEva				// Base evasion
	FlatEva				// Flat evasion
	EvaPer				// Evasion percent
	BaseCR				// Base critical rate
	FlatCR				// Flat critical rate
	CRPer				// Critical rate percent
	BaseCRRes			// Base critical rate resistance
	FlatCRRes			// Flat Critical rate resistance
	CRResPer			// Critical rate resistance percent
	BaseCD				// Base critical damage
	FlatCD				// Flat critical damage
	CDPer				// Critical damage percent
	BaseCDRes			// Base critical damage resistance
	FlatCDRes			// Flat critical damage resistance
	CDResPer			// Critical damage resistance percent
	BaseStab			// Base stability
	FlatStab			// Flat stability
	StabPer				// Stability percent
	BaseStabRate		// Base stability rate
	FlatStabRate		// Flat stability rate
	StabRatePer			// Stability rate percent
	BaseCCPwr			// Base crowd control power
	FlatCCPwr			// Flat crowd control power
	CCPwrPer			// Crowd control power percent
	BaseCCRes			// Base crowd control resistance
	FlatCCRes			// Flat crowd control resistance
	CCResPer			// Crowd control resistance percent
	BaseRecovBst		// Base recovery boost
	FlatRecovBst		// Flat recovery boost
	RecovBstPer			// Recovery boost percent
	BaseCostRecov		// Base cost recovery
	FlatCostRecov		// Flat cost recovery
	CostRecovPer		// Cost recovery percent
	BaseAtkSpd			// Base attack speed
	FlatAtkSpd			// Flat attack speed
	AtkSpdPer			// Attack speed percent
	BaseMoveSpd			// Base move speed
	FlatMoveSpd			// Flat move speed
	MoveSpdPer			// Move speed percent
	BaseBlkRateBonus	// Base block rate bonus
	FlatBlkRateBonus	// Flat block rate bonus
	BlkRateBonusPer		// Block rate bonus percent
	BaseDefPier			// Base defense piercing
	FlatDefPier			// Flat defense piercing
	DefPierPer			// Defense piercing percent
	BaseDmgDlt			// Base damage dealt
	FlatDmgDlt			// Flat damage dealt
	DmgDltPer			// Damage dealt percent
	BaseDmgRes			// Base damage resistance
	FlatDmgRes			// Flat damage resistance
	DmgResPer			// Damage resistance percent
	BaseExDmgDlt		// Base EX-special skill damage dealt
	FlatExDmgDlt		// Flat EX-special skill damage dealt
	ExDmgDltPer			// EX-special skill damage dealt percent
	BaseExploEff		// Base explosive effectiveness
	FlatExploEff		// Flat explosive effectiveness
	ExploEffPer			// Explosive effectiveness percent
	BasePierEff			// Base piercing effectiveness
	FlatPierEff			// Flat piercing effectiveness
	PierEffPer			// Piercing effectivenss percent
	BaseMystEff			// Base mystic effectiveness
	FlatMystEff			// Flat mystic effectiveness
	MystEffPer			// Mystic effectiveness percent
	BaseSonicEff		// Base sonic effectiveness
	FlatSonicEff		// Flat sonic effectiveness
	SonicEffPer			// Sonic effectiveness percent
	// Not sure what to do with "Siege" and "Normal" attack types as there are no modifiers to these types
	// Leave them commented out for now
	// BaseSiegeEff		// Base siege effectiveness
	// FlatSiegeEff		// Flat siege effectiveness
	// SiegeEffPer		// Siege effectiveness percent
	// BaseNormEff		// Base normal effectiveness
	// FlatNormEff		// Flat normal effectiveness
	// NormEffPer		// Normal effectiveness percent
	BaseBuffRent		// Base buff rentention
	FlatBuffRent		// Flat buff rentention
	BuffRentPer			// Buff rentention percent
	BaseDebuffRent		// Base debuff rentention
	FlatDebuffRent		// Flat debuff rentention
	DebuffRentPer		// Debuff rentention percent
	// For bosses, include groggy stat information
	GrogGauge			// Groggy gauge
	GrogDur				// Groggy duration
	StatsLen			// Delimiter for quickly making container sizes
)

var StatString = [...]string {
	"n/a",
	"base_hp",				// Base HP
	"flat_hp",				// Flat HP
	"hp%",					// HP percent
	"base_atk",				// Base attack
	"flat_atk",				// Flat attack
	"atk%",					// Attack percent
	"base_def",				// Base defense
	"flat_def",				// Flat defense
	"def%",					// Defense percent
	"base_heal",			// Base healing
	"flat_heal",			// Flat healing
	"heal%",				// Healing percent
	"base_acc",				// Base accuracy
	"flat_acc",				// Flat accuracy
	"acc%",					// Accuracy percent
	"base_eva",				// Base evasion
	"flat_eva",				// Flat evasion
	"eva%",					// Evasion percent
	"base_cr",				// Base critical rate
	"flat_cr",				// Flat critical rate
	"cr%",					// Critical rate percent
	"base_cr_res",			// Base critical rate resistance
	"flat_cr_res",			// Flat Critical rate resistance
	"cr_res%",				// Critical rate resistance percent
	"base_cd",				// Base critical damage
	"flat_cd",				// Flat critical damage
	"cd%",					// Critical damage percent
	"base_cd_res",			// Base critical damage resistance
	"flat_cd_res",			// Flat critical damage resistance
	"cd_res%",				// Critical damage resistance percent
	"base_stab",			// Base stability
	"flat_stab",			// Flat stability
	"stab%",				// Stability percent
	"base_stab_rate",		// Base stability rate
	"flat_stab_rate",		// Flat stability rate
	"stab_rate%",			// Stability rate percent
	"base_cc_pwr",			// Base crowd control power
	"flat_cc_pwr",			// Flat crowd control power
	"cc_pwr%",				// Crowd control power percent
	"base_cc_res",			// Base crowd control resistance
	"flat_cc_res",			// Flat crowd control resistance
	"cc_res%",				// Crowd control resistance percent
	"base_recov_bst",		// Base recovery boost
	"flat_recov_bst",		// Flat recovery boost
	"recov_bst%",			// Recovery boost percent
	"base_cost_recov",		// Base cost recovery
	"flat_cost_recov",		// Flat cost recovery
	"cost_recov%",			// Cost recovery percent
	"base_atk_spd",			// Base attack speed
	"flat_atk_spd",			// Flat attack speed
	"atk_spd%",				// Attack speed percent
	"base_move_spd",		// Base move speed
	"flat_move_spd",		// Flat move speed
	"move_spd%",			// Move speed percent
	"base_blk_rate_bonus",	// Base block rate bonus
	"flat_blk_rate_bonus",	// Flat block rate bonus
	"blk_rate_bonus%",		// Block rate bonus percent
	"base_def_pier",		// Base defense piercing
	"flat_def_pier",		// Flat defense piercing
	"def_pier%",			// Defense piercing percent
	"base_dmg_dlt",			// Base damage dealt
	"flat_dmg_dlt",			// Flat damage dealt
	"dmg_dlt%",				// Damage dealt percent
	"base_dmg_res",			// Base damage resistance
	"flat_Dmg_res",			// Flat damage resistance
	"dmg_res%",				// Damage resistance percent
	"base_ex_dmg_dlt",		// Base EX-special skill damage dealt
	"flat_ex_dmg_dlt",		// Flat EX-special skill damage dealt
	"ex_dmg_dlt%",			// EX-special skill damage dealt percent
	"base_explo_eff",		// Base explosive effectiveness
	"flat_explo_eff",		// Flat explosive effectiveness
	"explo_eff%",			// Explosive effectiveness percent
	"base_pier_eff",		// Base piercing effectiveness
	"flat_pier_eff",		// Flat piercing effectiveness
	"pier_eff%",			// Piercing effectiveness percent
	"base_myst_eff",		// Base mystic effectiveness
	"flat_myst_eff",		// Flat mystic effectiveness
	"myst_eff%",			// Mystic effectiveness percent
	"base_sonic_eff",		// Base sonic effectiveness
	"flat_sonic_eff",		// Flat sonic effectiveness
	"sonic_eff%",			// Sonic effectiveness percent
	// Not sure what to do with "Siege" and "Normal" attack types as there are no modifiers to these types
	// Leave them commented out for now
	// "base_siege_eff",	// Base siege effectiveness
	// "flat_siege_eff",	// Flat siege effectiveness
	// "siege_eff%",		// Siege effectiveness percent
	// "base_norm_eff",		// Base normal effectiveness
	// "flat_norm_eff",		// Flat normal effectiveness
	// "norm_eff%",			// Normal effectiveness percent
	"base_buff_rent",		// Base buff rentention
	"flat_buff_rent",		// Flat buff rentention
	"buff_rent%",			// Buff rentention percent
	"base_debuff_rent",		// Base debuff rentention
	"flat_debuff_rent",		// Flat debuff rentention
	"debuff_rent%",			// Debuff rentention percent
	// For bosses, include groggy stat information
	"grog_gauge",			// Groggy gauge
	"grod_dur",				// Groggy duration
	// Not necessary to have string equivalent for "StatsLen", but created one just to be safe
	"",
}