package modifer

import (

)

type ModType int

const (
	NoMod ModType = iota
	// Buffs
	AtkUp				// Attack up
	AtkSpdUp			// Attack speed up
	BlkRateUp			// Block rate up
	BulletTypeChange	// Bullet type change
	BulletTypeDmgUp		// Bullet type damage up
	CostDown			// Cost decrease
	AddCostDown			// Additional cost decrease
	CritRateUp			// Critical rate up
	CritResistUp		// Critical resistance up
	CritDmgUp			// Critical damage up
	CritDmgResistUp		// Critical damage resistance up
	IncomingDmgDown		// Incoming damage down
	DefUp				// Defense up
	EvaUp				// Evasion up
	ConRecovUp			// Continuous recovery up
	RecovBstUp			// Recovery boost up
	HealingUp			// Healing up
	AccUp				// Accuracy up
	MaxHpUp				// Max hp up
	SpdUp				// Speed up
	CCResistUp			// Crowd control resistance up
	CCPwrUp				// Crowd control power up
	RngUp				// Range up
	StabUp				// Stability up
	ViewUp				// View up
	Barrier				// Barrier
	EnhanceExplo		// Enhance explosion damage
	EnhancePier			// Enhance piercing damage
	EnhanceMyst			// Enhance mystic damage
	EnhanceSonic		// Enhance sonic damage
	EnhanceLightArm		// Enhance light armor
	EnhanceHeavyArm		// Enhance heavy armor
	EnhanceSpecArm		// Enhance special armor
	EnhanceElasArm		// Enhance elastic armor
	BuffDurUp			// Buff duration up
	CCDurUp				// Crowd control duration up
	DebuffDurUp			// Debuff duration up
	DmgLimit			// Damage limit up
	HealWhenHit			// Heal when hit
	FirstAid			// Recover hp at certain intervals
	AmmoUp				// Ammo up
	CostRegenUp			// Cost regeneration up
	CostOverload		// Cost overload
	DmgDealtUp			// Damage dealt up
	ExDmgDealtUp		// EX skill damage dealt up
	IncomingExDmgDown	// Incoming EX skill damage down
	// Debuffs
	AtkDown				// Attack down
	AtkSpdDown			// Attack speed down
	BlkRateDown			// Block rate down
	BulletTypeDmgDown	// Bullet type damage down
	Bleed				// Bleed
	Burn				// Burn
	BurnDmgUp			// Burn damage up
	Chill				// Chill
	ChillDmgUp			// Chill damage up
	CostUp				// Cost increase
	CostRegenDown		// Cost regeneration down
	CritRateDown		// Critical rate down
	CritResistDown		// Critical resistance down
	CritDmgDown			// Critical damage down
	CritDmgResistDown	// Critical damage resistance down
	Weakness			// Weakness
	IncomingDmgUp		// incoming damage up
	DefDown				// Defense down
	EvaDown				// Evasion down
	ConRecovDown		// Continuous recovery down
	RecovBstDown		// Recovery boost down
	HealingDown			// Healing down
	Curse				// Curse
	AccDown				// Accuracy down
	MaxHpDown			// Max hp down
	SpdDown				// Speed down
	CCResistDown		// Crowd control resistance down
	Poison				// Poison
	PoisonDmgUp			// Poison damage up
	RngDown				// Range down
	StabDown			// Stability down
	Unconscious			// ??? (what kind of debuff is this???)
	ViewDown			// View down
	DecreaseExplo		// Decrease explosion damage
	DecreasePier		// Decrease piercing damage
	DecreaseMyst		// Decrease mystic damage
	DecreaseSonic		// Decrease sonic damage
	DecreaseLightArm	// Decrease light armor
	DecreaseHeavyArm	// Decrease heavy armor
	DecreaseSpecArm		// Decrease special armor
	DecreaseElasArm		// Decrease elastic armor
	HealingLimitDown	// Healing limit down
	Cheer				// Cheer
	ElecShock			// Electric shock
	ElecShockDmgUp		// Eelctric shock damage up
	ExDmgDealtDown		// EX skill damage dealt down
	IncomingExDmgUp		// Incoming EX skill damage up
	// Both
	FocusedAssault		// Focused assault
	DefPen				// Defense peneration
	// Crowd control
	// Special
)