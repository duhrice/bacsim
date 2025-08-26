package stat

import (
	"math"
)

func (s Stat) String() string {
	return StatString[s]
}

func StrToStat(s string) Stat {
	for i, v := range StatString {
		// Check if stat type is valid
		if v == s {
			// Found valid stat type
			return Stat(i)
		}
	}
	// Invalid stat type found
	return -1
}

// Some calculations may use up to 2 decimal points of precision
// Uncomment this function when needed
// func roundFloat(val float64, precision uint) float64 {
// 	ratio := math.Pow(10, float64(precision))
// 	return math.Round(val*ratio) / ratio
// }

// Caluclations for final stats
// For now, only include function calculations for flat numbers (integers)
// Come back to this if functions calculations are needed for percent numbers
func (s Stats) TotalHp() float64 {
	return math.Round(s[BaseHp] + s[FlatHp]) * (1 + s[HpPer])
}
func (s Stats) TotalAtk() float64 {
	return math.Round(s[BaseAtk] + s[FlatAtk]) * (1 + s[AtkPer])
}
func (s Stats) TotalDef() float64 {
	return math.Round(s[BaseDef] + s[FlatDef]) * (1 + s[DefPer])
}
func (s Stats) TotalHeal() float64 {
	return math.Round(s[BaseHeal] + s[FlatHeal]) * (1 + s[HealPer])
}
func (s Stats) TotalAcc() float64 {
	return math.Round(s[BaseAcc] + s[FlatAcc]) * (1 + s[AccPer])
}
func (s Stats) TotalEva() float64 {
	return math.Round(s[BaseEva] + s[FlatEva]) * (1 + s[EvaPer])
}
func (s Stats) TotalCrit() float64 {
	return math.Round(s[BaseCR] + s[FlatCR]) * (1 + s[CRPer])
}
func (s Stats) TotalCritRes() float64 {
	return math.Round(s[BaseCRRes] + s[FlatCRRes]) * (1 + s[CRResPer])
}
// Critical damage is a percent number. Uncomment this if needed
// func (s Stats) TotalCritDmg() float64 {
//	return roundFloat((s[BaseCD] + (s[FlatCD]/10000)) * (1 + s[CDPer]), 2)
// }
// Critical damage resistance is a percent number. Uncomment this if needed
// func (s Stats) TotalCritDmgRes() float64 {
//	return roundFloat((s[BaseCDRes] + (s[FlatCDRes]/10000)) * (1 + s[CDResPer]), 2)
// }
func (s Stats) TotalStab() float64 {
	return math.Round(s[BaseStab] + s[FlatStab]) * (1 + s[StabPer])
}
// Stability rate is a percent number. Uncomment this if needed
// func (s Stats) TotalStabRate() float64 {
//	return roundFloat((s[BaseStabRate] + (s[FlatStabRate]/10000)) * (1 + s[StabRatePer]), 2)
// }
func (s Stats) TotalCCPwr() float64 {
	return math.Round(s[BaseCCPwr] + s[FlatCCPwr]) * (1 + s[CCPwrPer])
}
func (s Stats) TotalCCRes() float64 {
	return math.Round(s[BaseCCRes] + s[FlatCCRes]) * (1 + s[CCResPer])
}
func (s Stats) TotalRecovBst() float64 {
	return math.Round(s[BaseRecovBst] + s[FlatRecovBst]) * (1 + s[RecovBstPer])
}
func (s Stats) TotalCostRecov() float64 {
	return math.Round(s[BaseCostRecov] + s[FlatCostRecov]) * (1 + s[CostRecovPer])
}
// Attack speed is a percent number. Uncomment this if needed
// func (s Stats) TotalAtkSpd() float64 {
//	return roundFloat((s[BaseAtkSpd] + (s[FlatAtkSpd]/10000)) * (1 + s[AtkSpdPer]), 2)
// }
func (s Stats) TotalMovSpd() float64 {
	return math.Round(s[BaseMoveSpd] + s[FlatMoveSpd]) * (1 + s[MoveSpdPer])
}
// Block rate bonus is a percent number. Uncomment this if needed
// func (s Stats) TotalBlkRateBonus() float64 {
//	return roundFloat((s[BaseBlkRateBonus] + (s[FlatBlkRateBonus]/10000)) * (1 + s[BlkRateBonusPer]), 2)
// }
func (s Stats) TotalDefPier() float64 {
	return math.Round(s[BaseDefPier] + s[FlatDefPier]) * (1 + s[DefPierPer])
}
// Damage dealt is a percent number. Uncomment this if needed
// func (s Stats) TotalDmgDlt() float64 {
//	return roundFloat((s[BaseDmgDlt] + (s[FlatDmgDlt]/10000)) * (1 + s[DmgDltPer]), 2)
// }
// Damage resist is a percent number. Uncomment this if needed
// func (s Stats) TotalDmgRes() float64 {
//	return roundFloat((s[BaseDmgRes] + (s[FlatDmgRes]/10000)) * (1 + s[DmgResPer]), 2)
// }
// Ex special skill damage dealt is a percent number. Uncomment this if needed
// func (s Stats) TotalExDmgDlt() float64 {
//	return roundFloat((s[BaseExDmgDlt] + (s[FlatExDmgDlt]/10000)) * (1 + s[ExDmgDltPer]), 2)
// }
// Explosive effectiveness is a percent number. Uncomment this if needed
// func (s Stats) TotalExploEff() float64 {
//	return roundFloat((s[BaseExploEff] + (s[FlatExploEff]/10000)) * (1 + s[ExploEffPer]), 2)
// }
// Piercing effectiveness is a percent number. Uncomment this if needed
// func (s Stats) TotalPierEff() float64 {
//	return roundFloat((s[BasePierEff] + (s[FlatPierEff]/10000)) * (1 + s[PierEffPer]), 2)
// }
// Mystic effectiveness is a percent number. Uncomment this if needed
// func (s Stats) TotalMystEff() float64 {
//	return roundFloat((s[BaseMystEff] + (s[FlatMystEff]/10000)) * (1 + s[MystEffPer]), 2)
// }
// Sonic effectiveness is a percent number. Uncomment this if needed
// func (s Stats) TotalSonicEff() float64 {
//	return roundFloat((s[BaseSonicEff] + (s[FlatSonicEff]/10000)) * (1 + s[SonicEffPer]), 2)
// }
// Buff rentention is a percent number. Uncomment this if needed
// func (s Stats) TotalBuffRent() float64 {
//	return roundFloat((s[BaseBuffRent] + (s[FlatBuffRent]/10000)) * (1 + s[BuffRentPer]), 2)
// }
// Debuff rentention is a percent number. Uncomment this if needed
// func (s Stats) TotalDebuffRent() float64 {
//	return roundFloat((s[BaseDebuffRent] + (s[FlatDebuffRent]/10000)) * (1 + s[DebuffRentPer]), 2)
// }