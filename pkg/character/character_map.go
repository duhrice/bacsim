package character

import (
	"github.com/duhrice/bacsim/pkg/affiliation"
	"github.com/duhrice/bacsim/pkg/class"
	"github.com/duhrice/bacsim/pkg/role"
	"github.com/duhrice/bacsim/pkg/position"
	"github.com/duhrice/bacsim/pkg/attack"
	"github.com/duhrice/bacsim/pkg/defense"
	"github.com/duhrice/bacsim/pkg/mood"
)

// All students have three terrain moods (city/urban/street; desert/field/outdoor; indoor)
type CombatPower struct {
	Urban mood.Mood
	Field mood.Mood
	Indoor mood.Mood
}

type CharInfo struct {
	ID Char
	Name string
	School affiliation.School
	Club affiliation.Club
	BaseStarGrade int
	Class class.ClassType
	Role role.RoleType
	// summons 		<- implementation on hold for now
	Position position.PositionType
	Attack attack.AttackType
	Defense defense.DefenseType
	TerrainMoods CombatPower
}

// Weapons give bonus combat power for terrains. Use base combat power for now and update it during student setup
var CharInfoMap = map[Char]CharInfo {
	TestChar: {
		ID: TestChar,
		Name: TestChar.String(),
		School: affiliation.Unknown,
		Club: affiliation.Club(affiliation.Unknown),
		BaseStarGrade: 1,
		Class: class.Unknown,
		Role: role.Unknown,
		// summons		<- implementation on hold for now
		Position: position.Unknown,
		Attack: attack.Unknown,
		Defense: defense.Unknown,
		TerrainMoods: CombatPower {
			Urban: mood.Unknown,
			Field: mood.Unknown,
			Indoor: mood.Unknown,
		},		
	},
	Aru: {
		ID: Aru,
		Name: Aru.String(),
		School: affiliation.Gehenna,
		Club: affiliation.ProblemSolver68,
		BaseStarGrade: 3,
		Class: class.Striker,
		Role: role.Dealer,
		// summons		<- implementation on hold for now
		Position: position.Back,
		Attack: attack.Explosive,
		Defense: defense.Light,
		TerrainMoods: CombatPower {
			Urban: mood.S,
			Field: mood.B,
			Indoor: mood.D,
		},
	},
}