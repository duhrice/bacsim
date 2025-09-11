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
	FocusedAssault		// Focused assault
	ConRecovUp			// Continuous recovery up
	RecovBstUp			// Recovery boost up
	HealingUp			// Healing up
	AccUp				// Accuracy up
	MaxHpUp				// MaxHpUp
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
	DefPen				// Defense peneration
	ExDmgDealtUp		// EX skill damage dealt up
	IncomingExDmgDown	// Incoming EX skill damage down
	// Debuffs
	// Crowd control
	// Special
)