package attack

import (

)

type AttackType int

// "Siege" is a special attack type but no one uses it
// "Siege" is effective against in-game cover which has the "Structure" defense type
const (
	Explosive AttackType = iota
	Piercing
	Mystic
	Sonic
	Siege
	Normal
	Unknown
)

var AttackTypeString = [...]string {
	"explosive",
	"piercing",
	"mystic",
	"sonic",
	"siege",
	"normal",
	"unknown",
}