package modifier

import (

)

type ModType int

const (
	NoMod ModType = iota
	// Buffs
	AtkUp					// Attack up
	AtkSpdUp				// Attack speed up
	BlkRateUp				// Block rate up
	BulletTypeChange		// Bullet type change
	BulletTypeDmgUp			// Bullet type damage up
	CostDown				// Cost decrease
	AddCostDown				// Additional cost decrease
	CritRateUp				// Critical rate up
	CritResistUp			// Critical resistance up
	CritDmgUp				// Critical damage up
	CritDmgResistUp			// Critical damage resistance up
	IncomingDmgDown			// Incoming damage down
	DefUp					// Defense up
	EvaUp					// Evasion up
	ConRecovUp				// Continuous recovery up
	RecovBstUp				// Recovery boost up
	HealingUp				// Healing up
	AccUp					// Accuracy up
	MaxHpUp					// Max hp up
	SpdUp					// Speed up
	CCResistUp				// Crowd control resistance up
	CCPwrUp					// Crowd control power up
	RngUp					// Range up
	StabUp					// Stability up
	ViewUp					// View up
	Barrier					// Barrier
	EnhanceExplo			// Enhance explosion damage
	EnhancePier				// Enhance piercing damage
	EnhanceMyst				// Enhance mystic damage
	EnhanceSonic			// Enhance sonic damage
	EnhanceLightArm			// Enhance light armor
	EnhanceHeavyArm			// Enhance heavy armor
	EnhanceSpecArm			// Enhance special armor
	EnhanceElasArm			// Enhance elastic armor
	BuffDurUp				// Buff duration up
	CCDurUp					// Crowd control duration up
	DebuffDurUp				// Debuff duration up
	DmgLimit				// Damage limit up
	HealWhenHit				// Heal when hit
	FirstAid				// Recover hp at certain intervals
	AmmoUp					// Ammo up
	CostRegenUp				// Cost regeneration up
	CostOverload			// Cost overload
	DmgDealtUp				// Damage dealt up
	ExDmgDealtUp			// EX skill damage dealt up
	IncomingExDmgDown		// Incoming EX skill damage down
	// Debuffs
	AtkDown					// Attack down
	AtkSpdDown				// Attack speed down
	BlkRateDown				// Block rate down
	BulletTypeDmgDown		// Bullet type damage down
	Bleed					// Bleed
	Burn					// Burn
	BurnDmgUp				// Burn damage up
	Chill					// Chill
	ChillDmgUp				// Chill damage up
	CostUp					// Cost increase
	CostRegenDown			// Cost regeneration down
	CritRateDown			// Critical rate down
	CritResistDown			// Critical resistance down
	CritDmgDown				// Critical damage down
	CritDmgResistDown		// Critical damage resistance down
	Weakness				// Weakness
	IncomingDmgUp			// incoming damage up
	DefDown					// Defense down
	EvaDown					// Evasion down
	ConRecovDown			// Continuous recovery down
	RecovBstDown			// Recovery boost down
	HealingDown				// Healing down
	Curse					// Curse
	AccDown					// Accuracy down
	MaxHpDown				// Max hp down
	SpdDown					// Speed down
	CCResistDown			// Crowd control resistance down
	Poison					// Poison
	PoisonDmgUp				// Poison damage up
	RngDown					// Range down
	StabDown				// Stability down
	Unconscious				// ???*
	ViewDown				// View down
	DecreaseExplo			// Decrease explosion damage
	DecreasePier			// Decrease piercing damage
	DecreaseMyst			// Decrease mystic damage
	DecreaseSonic			// Decrease sonic damage
	DecreaseLightArm		// Decrease light armor
	DecreaseHeavyArm		// Decrease heavy armor
	DecreaseSpecArm			// Decrease special armor
	DecreaseElasArm			// Decrease elastic armor
	HealingLimitDown		// Healing limit down
	Cheer					// Cheer
	ElecShock				// Electric shock
	ElecShockDmgUp			// Eelctric shock damage up
	ExDmgDealtDown			// EX skill damage dealt down
	IncomingExDmgUp			// Incoming EX skill damage up
	// Both
	FocusedAssault			// Focused assault
	DefPen					// Defense peneration
	// Crowd control
	Blind					// Blind
	Charm					// Charm
	Confusion				// Confusion
	EMP						// EMP
	Fear					// Fear
	Inoperative				// ???*
	MindControl				// Mind control
	Paralysis				// Paralysis
	Provoke					// Provoke
	Purify					// ???*
	Silence					// Silence
	Freeze					// Freeze
	Stun					// Stun
	// Special
	AccumDmg				// Accumulate damage
	EnChargeN				// Energy charge none
	EnChargeH				// Energy charge half
	EnChargeF				// Energy charge full
	FormChange				// Form change
	Fury					// Fury
	Immortal				// Immortal
	CCImmune				// Crowd control immunization
	DmgImmune				// Damage immunization
	Imp						// Imp
	EvilDeed				// Evil deed
	StampsOfPraise			// Stamps of praise
	HolyNightBless			// Holy night blessing
	SantaGift				// ???*
	ProbExAddDmg			// Probability EX skill additional damage
	SilverBullet			// Silver bullet
	AbiEshuh				// Abi-Eshuh
	Omamori					// Omamori
	Serenity				// Serenity
	SkillBulletTypeChange	// ???*
	CostRegenDisable		// Cost regeneration disable
	Drill					// Drill
	Out						// Out
	Prayer					// Prayer
	Zealous					// Zealous
	Fuurinkazan				// Fuurinkazan
	RoastedPotato			// Roasted potato
	SuccessfulProtect		// Successful protection
	Infiltration			// Infilitration
	SugarRush				// Sugar rush
	Clue					// Clue
	EnhanceExDmg			// Enhance EX skill damage
	Bulletproof				// Bullet proof
	SuppressiveFire			// Suppressive fire
	RetreatGrace			// Retreat grace
	RetreatDeferral			// Retreat deferral
	Flyer					// Flyer
	AddPoisonDmg			// Additional poison damage
	ChillReducePeriod		// Chill reduce period
	Fanservice				// Fanservice
	Penlights				// Penlights
	FavoritePillow			// Favorite pillow
	ExDrawStock				// Ex skill draw stock
	Duel					// Duel
	Spirit					// Spirit
	SpiritF					// Spirit full
	ExSkillCardDup			// Ex skill card duplication
	BattleSense				// Battle sense
	BattleSenseF			// Battle sense full
	// All student special effects up to week of September 8, 2025 (Global). More effects to add in the future
	// Below are some boss/raid specific special effects
	DebuffCountGreen		// Debuff count green
	DebuffCountYellow		// Debuff count yellow
	DebuffCountRed			// Debuff count red
	Obelisk					// Obelisk
	BallistaFiring			// Ballista firing
	UnstableFlame			// Unstable flame
	ModLen					// Delimiter for quickly making container sizes
)
// * what kind of debuffs are these???

var ModTypeString = [...]string {
	"n/a",
	// Buffs
	"atk_up",				// Attack up
	"atk_spd_up",			// Attack speed up
	"blk_rate_up",			// Block rate up
	"bullet_type_change",	// Bullet type change
	"bullet_type_dmg_up",	// Bullet type damage up
	"cost_down",			// Cost decrease
	"add_cost_down",		// Additional cost decrease
	"crit_rate_up",			// Critical rate up
	"crit_resist_up",		// Critical resistance up
	"crit_dmg_up",			// Critical damage up
	"crit_dmg_resist_up",	// Critical damage resistance up
	"incoming_dmg_down",	// Incoming damage down
	"def_up",				// Defense up
	"eva_up",				// Evasion up
	"con_recov_up",			// Continuous recovery up
	"recov_bst_up",			// Recovery boost up
	"healing_up",			// Healing up
	"acc_up",				// Accuracy up
	"max_hp_up",			// Max hp up
	"spd_up",				// Speed up
	"cc_resist_up",			// Crowd control resistance up
	"cc_pwr_up",			// Crowd control power up
	"rng_up",				// Range up
	"stab_up",				// Stability up
	"view_up",				// View up
	"barrier",				// Barrier
	"enhance_explo",		// Enhance explosion damage
	"enhance_pier",			// Enhance piercing damage
	"enhance_myst",			// Enhance mystic damage
	"enhance_sonic",		// Enhance sonic damage
	"enhance_light_arm",	// Enhance light armor
	"enhance_heavy_arm",	// Enhance heavy armor
	"enhance_spec_arm",		// Enhance special armor
	"enhance_elas_arm",		// Enhance elastic armor
	"buff_dur_up",			// Buff duration up
	"cc_dur_up",			// Crowd control duration up
	"debuff_dur_up",		// Debuff duration up
	"dmg_limit",			// Damage limit up
	"heal_when_hit",		// Heal when hit
	"first_aid",			// Recover hp at certain intervals
	"ammo_up",				// Ammo up
	"cost_regen_up",		// Cost regeneration up
	"cost_overload",		// Cost overload
	"dmg_dealt_up",			// Damage dealt up
	"ex_dmg_dealt_up",		// EX skill damage dealt up
	"incoming_ex_dmg_down",	// Incoming EX skill damage down
	// Debuffs
	AtkDown					// Attack down
	AtkSpdDown				// Attack speed down
	BlkRateDown				// Block rate down
	BulletTypeDmgDown		// Bullet type damage down
	Bleed					// Bleed
	Burn					// Burn
	BurnDmgUp				// Burn damage up
	Chill					// Chill
	ChillDmgUp				// Chill damage up
	CostUp					// Cost increase
	CostRegenDown			// Cost regeneration down
	CritRateDown			// Critical rate down
	CritResistDown			// Critical resistance down
	CritDmgDown				// Critical damage down
	CritDmgResistDown		// Critical damage resistance down
	Weakness				// Weakness
	IncomingDmgUp			// incoming damage up
	DefDown					// Defense down
	EvaDown					// Evasion down
	ConRecovDown			// Continuous recovery down
	RecovBstDown			// Recovery boost down
	HealingDown				// Healing down
	Curse					// Curse
	AccDown					// Accuracy down
	MaxHpDown				// Max hp down
	SpdDown					// Speed down
	CCResistDown			// Crowd control resistance down
	Poison					// Poison
	PoisonDmgUp				// Poison damage up
	RngDown					// Range down
	StabDown				// Stability down
	Unconscious				// ???*
	ViewDown				// View down
	DecreaseExplo			// Decrease explosion damage
	DecreasePier			// Decrease piercing damage
	DecreaseMyst			// Decrease mystic damage
	DecreaseSonic			// Decrease sonic damage
	DecreaseLightArm		// Decrease light armor
	DecreaseHeavyArm		// Decrease heavy armor
	DecreaseSpecArm			// Decrease special armor
	DecreaseElasArm			// Decrease elastic armor
	HealingLimitDown		// Healing limit down
	Cheer					// Cheer
	ElecShock				// Electric shock
	ElecShockDmgUp			// Eelctric shock damage up
	ExDmgDealtDown			// EX skill damage dealt down
	IncomingExDmgUp			// Incoming EX skill damage up
	// Both
	"focused_assault",		// Focused assault
	"def_pen",				// Defense peneration
	// Crowd control
	"blind",				// Blind
	"charm",				// Charm
	"confusion",			// Confusion
	"emp",					// EMP
	"fear",					// Fear
	"inoperative",			// ???*
	"mind_control",			// Mind control
	"paralysis",			// Paralysis
	"provoke",				// Provoke
	"purify",				// ???*
	"silence",				// Silence
	"freeze",				// Freeze
	"stun",					// Stun
	// Special
	AccumDmg				// Accumulate damage
	EnChargeN				// Energy charge none
	EnChargeH				// Energy charge half
	EnChargeF				// Energy charge full
	FormChange				// Form change
	Fury					// Fury
	Immortal				// Immortal
	CCImmune				// Crowd control immunization
	DmgImmune				// Damage immunization
	Imp						// Imp
	EvilDeed				// Evil deed
	StampsOfPraise			// Stamps of praise
	HolyNightBless			// Holy night blessing
	SantaGift				// ???*
	ProbExAddDmg			// Probability EX skill additional damage
	SilverBullet			// Silver bullet
	AbiEshuh				// Abi-Eshuh
	Omamori					// Omamori
	Serenity				// Serenity
	SkillBulletTypeChange	// ???*
	CostRegenDisable		// Cost regeneration disable
	Drill					// Drill
	Out						// Out
	Prayer					// Prayer
	Zealous					// Zealous
	Fuurinkazan				// Fuurinkazan
	RoastedPotato			// Roasted potato
	SuccessfulProtect		// Successful protection
	Infiltration			// Infilitration
	SugarRush				// Sugar rush
	Clue					// Clue
	EnhanceExDmg			// Enhance EX skill damage
	Bulletproof				// Bullet proof
	SuppressiveFire			// Suppressive fire
	RetreatGrace			// Retreat grace
	RetreatDeferral			// Retreat deferral
	Flyer					// Flyer
	AddPoisonDmg			// Additional poison damage
	ChillReducePeriod		// Chill reduce period
	Fanservice				// Fanservice
	Penlights				// Penlights
	FavoritePillow			// Favorite pillow
	ExDrawStock				// Ex skill draw stock
	Duel					// Duel
	Spirit					// Spirit
	SpiritF					// Spirit full
	ExSkillCardDup			// Ex skill card duplication
	BattleSense				// Battle sense
	BattleSenseF			// Battle sense full
	// All student special effects up to week of September 8, 2025 (Global). More effects to add in the future
	// Below are some boss/raid specific special effects
	DebuffCountGreen		// Debuff count green
	DebuffCountYellow		// Debuff count yellow
	DebuffCountRed			// Debuff count red
	Obelisk					// Obelisk
	BallistaFiring			// Ballista firing
	UnstableFlame			// Unstable flame
	ModLen					// Delimiter for quickly making container sizes
}