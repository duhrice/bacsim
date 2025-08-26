package defense

import (

)

type DefenseType int

// "Structure" is a special defense type only used by in-game cover
// "Structure" is weak to the "Siege" attack type
const (
	Light DefenseType = iota
	Heavy
	Special
	Elastic
	Structure
	Normal
	Unknown
)

var DefenseTypeString = [...]string {
	"light",
	"heavy",
	"special",
	"elastic",
	"structure",
	"normal",
	"unknown",
}