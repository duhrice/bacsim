package attributes

type (
	// Int representing student stat
	Stat int
	// Array of stat types
	Stats [EndStatType]float64
)

// All student and enemy stat types
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
	BaseNormAtKRange	// Base normal attack range
	FlatNormAtkRange	// Flat normal attack range
	NormAtkRangePer		// Normal attack range percent
	BaseCCPower			// Base crowd control power
	FlatCCPower			// Flat crowd control power
	CCPowerPer			// Crowd control power percent
	BaseCCRes			// Base crowd control resistance
	FlatCCRes			// Flat crowd control resistance
	CCResPer			// Crowd control resistance percent
	BaseRecovBoost		// Base recovery boost
	FlatRecovBoost		// Flat recovery boost
	RecovBoostPer		// Recovery boost percent

	// More to add...
)