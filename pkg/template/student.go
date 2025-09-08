package template

import (
	"github.com/duhrice/bacsim/pkg/character"
	"github.com/duhrice/bacsim/pkg/class"
	"github.com/duhrice/bacsim/pkg/role"
	"github.com/duhrice/bacsim/pkg/attack"
	"github.com/duhrice/bacsim/pkg/defense"
	"github.com/duhrice/bacsim/pkg/affiliation"
	"github.com/duhrice/bacsim/pkg/mood"
	"github.com/duhrice/bacsim/pkg/weapon"
	"github.com/duhrice/bacsim/pkg/position"
)

// What methods are shared by all students?
type Student interface {

}

// What parameters should all students have?
type StudentWrapper struct {
	// Basic information
	Index character.Char			// All students have an assigned index
	StudentStars int				// Student base rarity ranges from 1 to 3 stars. All students can have max of 5.
	StudentLevel int				// Current student level. Current max student level is 90 (will be 100 in future updates)
	WeaponStars int					// Weapon has 0 to 3 stars. All weapons have a max of 3.
	WeaponLevel int					// Current weapon level. Current weapon level is dependent on the number of weapon stars
	Class class.ClassType			// Striker or special class
	Role role.RoleType				// Team role (tank, damage dealer, healer, support, tactical support)
	Attack attack.AttackType		// Attack type (explosive, piercing, mystic, sonic) (no students use normal)
	Defense defense.DefenseType		// Defense type (light, heavy, special, elastic) (no students use normal)
	School affiliation.School		// Kivotos schools
	Club affiliation.Club			// School clubs
	TerrainMoods CombatPower		// All students have three terrain moods (city/urban/street; desert/field/outdoor; indoor)
	Weapon weapon.WeaponType		// Student weapon types
	Position position.PositionType	// Student positions in a team (front, middle, back)				

	Student							// All students should use the same methods, hence the interface
}

type CombatPower struct {
	Urban mood.Mood
	Field mood.Mood
	Indoor mood.Mood
}