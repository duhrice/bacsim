package template

import (
	"github.com/duhrice/bacsim/pkg/affiliation"
	"github.com/duhrice/bacsim/pkg/attack"
	"github.com/duhrice/bacsim/pkg/character"
	"github.com/duhrice/bacsim/pkg/class"
	"github.com/duhrice/bacsim/pkg/defense"
	"github.com/duhrice/bacsim/pkg/equipment"
	"github.com/duhrice/bacsim/pkg/mood"
	"github.com/duhrice/bacsim/pkg/position"
	"github.com/duhrice/bacsim/pkg/role"
	"github.com/duhrice/bacsim/pkg/weapon"
)

// What methods are shared by all students?
type Student interface {

}

// What parameters should all students have?
type StudentWrapper struct {
	// Basic student information
	Index character.Char			// All students have an assigned index
	StudentStars int				// Student base rarity ranges from 1 to 3 stars. All students can have max of 5.
	StudentLevel int				// Current student level. Current max student level is 90 (will be 100 in future updates)
	Class class.ClassType			// Striker or special class
	Role role.RoleType				// Team role (tank, damage dealer, healer, support, tactical support)
	Attack attack.AttackType		// Attack type (explosive, piercing, mystic, sonic) (no students use normal)
	Defense defense.DefenseType		// Defense type (light, heavy, special, elastic) (no students use normal)
	School affiliation.School		// Kivotos schools
	Club affiliation.Club			// School clubs
	TerrainMoods CombatPower		
	Position position.PositionType	// Student positions in a team (front, middle, back)
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

	Student							// All students should use the same methods, hence the interface
}

// All students have three terrain moods (city/urban/street; desert/field/outdoor; indoor)
type CombatPower struct {
	Urban mood.Mood
	Field mood.Mood
	Indoor mood.Mood
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
type Accessory struct {
	Tier int
	Level int
}