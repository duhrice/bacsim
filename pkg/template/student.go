package template

import (
	"github.com/duhrice/bacsim/pkg/character"
	"github.com/duhrice/bacsim/pkg/equipment"
	"github.com/duhrice/bacsim/pkg/weapon"
	"github.com/duhrice/bacsim/pkg/modifier"
	"github.com/duhrice/bacsim/pkg/source"
)

var DEFAULT_STAR_SCALING = [3][5]float32 {
	{1, 1.1, 1.22, 1.36, 1.53},			// ATK scaling
	{1, 1.05, 1.12, 1.21, 1.35},		// HP scaling
	{1, 1.075, 1.175, 1.295, 1.445},	// Healing scaling
}

const DEFAULT_STAT_GROWTH_TYPE string = "standard"

type fn func() error		// May need to modify this later to take "target" as an argument

// What methods are shared by all students?
type StudentHelper interface {

}

// What parameters should all students have?
type Student struct {
	// Basic student information
	StudentInfo character.CharInfo
	StudentStars int				// Student base rarity ranges from 1 to 3 stars. All students can have max of 5.
	StudentLevel int				// Current student level. Current max student level is 90 (will be 100 in future updates)
	BondLevel int					// Student bond level
	AddBondLevel int				// Additional bond levels for students with variants
	// Basic weapon infomation
	Weapon weapon.WeaponType		// Student weapon types
	WeaponStars int					// Weapon has 0 to 3 stars. All weapons have a max of 3.
	WeaponLevel int					// Current weapon level. Max weapon level is dependent on the number of weapon stars
	// Basic skill information
	Skills StudentSkills			
	// Student stats information
	Stats []float64					// Array containing all possible student stats
	// Student equipment information
	Equipment map[equipment.EquipmentType]Accessory
	// Student modifiers information
	Buffs []Modifier
	Debuffs []Modifier
	CrowdControl []Modifier
	Special	[]Modifier				// Includes unique effects, stacks, mechanics provided or used by students
	// Student actions information
	Action []fn						// Action that student is currently performing
	ActionQueue []fn				// Action queue for what student should perform next

	StudentHelper					// All students should use the same methods, hence the interface
}

// All students have four skills that can be leveled to a max of 10 (labelled as MAX in game)(ex, basic, enhanced, sub)
type StudentSkills struct {
	EX int
	Basic int
	Enhanced int
	Sub int
}

// All students can hace 3 equipments, with some students having an additional unique item
// First: hat, gloves, shoes
// Second: bag, badge, hairpin
// Third: charm, watch, necklace
// The accessory for each equipment slot is tiered from T0 - T10
// Each tier has levels from 1 to 10
// None of the above 2 lines applies for unique items which only has 2 tiers and no levels
type Accessory struct {
	Tier int
	Level int
}

// Modifiers are things that affect student stats and gameplay
// Two similar modifiers from the same source can not stack (i.e. two attack buffs from an ex-skill) except for special statuses
// Only the most recent modifier from a source is applied. The former will be overwritten
// These modifiers can be targeted at allies, self, enemies, or in AOE
type Modifier struct {
	ID modifier.ModType
	Source source.SourceType
	Value float32
	Duration int
}

// Constructor function for creating a new student
func NewStudent(
	name string,	
	level int,
	stars int,				// I think it is current student stars, not base student stars (I could be wrong though)
	scaling [][]float32,	// Student HP, ATK, and DEf scales based on star level
	growthType string,		// Some students may have unique stat growths
) (*Student, error) {
	s := &Student{

	}

	return s, nil
}