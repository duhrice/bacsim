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