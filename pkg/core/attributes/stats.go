package attributes

import (
	"strconv"
	"strings"
)

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
	BaseMagCount		// Base magazine count
	FlatMagCount		// Flat magazine count
	MagCountPer			// Magazine count percent
	NormAtkCost			// Normal attack cost
	BaseDmgDealt		// Base damage dealt
	FlatDmgDealt		// Flat damage dealt
	DmgDealtPer			// Damage dealt percent
	BaseExDmgDealt		// Base EX-special skill damage dealt
	FlatExDmgDealt		// Flat EX-special skill damage dealt
	ExDmgDealtPer		// EX-special skill damage dealt percent
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
	// SiegeEffPer			// Siege effectiveness percent
	// BaseNormEff			// Base normal effectiveness
	// FlatNormEff			// Flat normal effectiveness
	// NormEffPer			// Normal effectiveness percent
	BaseBuffRent		// Base buff rentention
	FlatBuffRent		// Flat buff rentention
	BuffRentPer			// Buff rentention percent
	BaseDebuffRent		// Base debuff rentention
	FlatDebuffRent		// Flat debuff rentention
	DebuffRentPer		// Debuff rentention percent
	// For bosses, include groggy stat information
	GrogGauge			// Groggy gauge
	GrogDur				// Groggy duration
	// Delimiter (not sure why this is needed, will have to do more research)
	// Most likely for quickly making slice sizes
	EndStatType
)

// Return stats as strings instead of ints
func (s Stat) String() string {
	return StatTypeString[s]
}

// Caluclations for final stats

// Print stats in a clear and readable manner
func PrettyPrintStats(stats []float64) string {
	// Create string builder for string format
	var sb strings.Builder
	// Gp through slice of stats to print
	for i, v := range stats {
		if v > 0 {
			// Get the corresponding stat
			sb.WriteString(StatTypeString[i])
			sb.WriteString(": ")
			// Get the value corresponding to the stat
			sb.WriteString(strconv.FormatFloat(v, 'f', 2, 32))
			sb.WriteString(" ")
		}
	}
	// Trim any empty spaces
	return strings.Trim(sb.String(), " ")
}

// Print stats slice in a clear and readable manner
func PrettyPrintStatsSlice(stats []float64) []string {
	// Make a slice to store final formatted stat slice
	r := make([]string, 0)
	// Create string builder for string format
	var sb strings.Builder
	// Go through slice of stats to print
	for i, v := range stats {
		if v == 0 {
			continue
		}
		// Get the corresponding stat
		sb.WriteString(StatTypeString[i])
		sb.WriteString(": ")
		// Get the value corresponding to the stat
		sb.WriteString(strconv.FormatFloat(v, 'f', 2, 32))
		// Add formmated string to new slice
		r = append(r, sb.String())
		// Clear string builder for next stat in slice
		sb.Reset()
	}
	// Return formatted slice of stats
	return r
}

var StatTypeString = [...]string {

}

// Convert stat type string to stat type int
func StrToStatType(s string) Stat {
	for i, v := range StatTypeString {
		// Check if stat type is valid
		if v == s {
			// Found valid stat type
			return Stat(i)
		}
	}
	// Invalid stat type found
	return -1
}