package template

import (
	"github.com/duhrice/bacsim/pkg/character"
	"github.com/duhrice/bacsim/pkg/class"
	"github.com/duhrice/bacsim/pkg/attack"
	"github.com/duhrice/bacsim/pkg/defense"
)

// What methods are shared by all students?
type Student interface {

}

// What parameters should all students have?
type StudentWrapper struct {
	// Basic information
	Index character.Char			// All students have an assigned index
	StudentStars int				// Student base rarity ranges from 1 to 3 stars. All students can have max of 5.
	StudentLevel int				// Current student level. Current max student level is 90 (will be 100 in future updates, edit this when it has been implemented)
	WeaponStars int					// Weapon has 0 to 3 stars. All weapons have a max of 3.
	WeaponLevel int					// Current weapon level. Current weapon level is dependent on the number of weapon stars
	Position class.ClassType		// Striker or special class
	// Add role here
	Attack attack.AttackType		// Attack type (explosive, piercing, mystic, sonic) (no students use normal)
	Defense defense.DefenseType		// Defense type (light, heavy, special, elastic) (no students use normal)
	// Add school and club here
	// Add mood here
	// Add equipment type and position here

	Student							// All students should use the same methods, hence the interface
}