package attributes

import (
	"encoding/json"
	"errors"
	"strings"
)

// Int representing student and enemy attack type
type AttackType int
// INt representing student and enemy defense type
type DefenseType int

// All student and enemy attack types
// "Siege" is a special attack type but no one uses it
// "Siege" is effective against in-game cover which has the "Structure" defense type
const (
	Explosive AttackType = iota
	Piercing
	Mystic
	Sonic
	Siege
	Normal_AT
	Unknown_AT
)

// All student and enemy defense types
// "Structure" is a special defense type but no one uses it
// "Structure" is used by in-game cover
const (
	Light DefenseType = iota
	Heavy
	Special
	Elastic
	Structure
	Normal_DT
	Unknown_DT
)

// Though about implementing a Go interface but thats already done and defined in other libraries
// Return attack types as strings instead of ints
func (at AttackType) String() string {
	return AttackTypeString[at]
}

// Return defense types as strings instead of ints
func (dt DefenseType) String() string {
	return DefenseTypeString[dt]
}

// JSON marshalling for attack type
func (at AttackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttackTypeString[at])
}

// JSON marshalling for attack type
func (dt DefenseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(DefenseTypeString[dt])
}

// JSON unmarshalling for attack type
func (at *AttackType) UnmarshalJSON(b []byte) error {
	// Create a variable to hold the unmarshalled JSON string
	var s string
	// Unmarshal JSON first, then error check
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	// Lower string for easier checking
	s = strings.ToLower(s)
	for i, v := range AttackTypeString {
		// Check if the attack type is a valid
		if v == s {
			// Found valid attack type
			*at = AttackType(i)
			return nil
		}
	}
	// Invalid attack type found
	return errors.New("unrecognized attack type")
}

// JSON unmarshalling for attack type
func (dt *DefenseType) UnmarshalJSON(b []byte) error {
	// Create a variable to hold the unmarshalled JSON string
	var s string
	// Unmarshal JSON first, then error check
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	// Lower string for easier checking
	s = strings.ToLower(s)
	for i, v := range DefenseTypeString {
		// Check if the attack type is a valid
		if v == s {
			// Found valid attack type
			*dt = DefenseType(i)
			return nil
		}
	}
	// Invalid attack type found
	return errors.New("unrecognized defense type")
}

// Attack types as strings
var AttackTypeString = [...]string {
	"explosive",
	"piercing",
	"mystic",
	"sonic",
	"siege",
	"normal_at",
	"unknown_at",
}

// Defense types as strings
var DefenseTypeString = [...]string {
	"light",
	"heavy",
	"special",
	"elastic",
	"structure",
	"normal_dt",
	"unknown_dt",
}

// Convert an attack type string to attack type int
func StringtoAT(s string) AttackType {
	for i, v := range AttackTypeString {
		// Check if attack type is valid
		if v == s {
			// Found valid attack type
			return AttackType(i)
		}
	}
	// Invalid attack type found
	return Unknown_AT
}

// Convert an attack type string to attack type int
func StringtoDT(s string) DefenseType {
	for i, v := range DefenseTypeString {
		// Check if attack type is valid
		if v == s {
			// Found valid attack type
			return DefenseType(i)
		}
	}
	// Invalid attack type found
	return Unknown_DT
}

// Convert attack type to attack type effectiveness percent
func AtkTypeToEff(at AttackType) Stat {
	switch at {
	case Explosive:
		return ExploEffPer
	case Piercing:
		return PierEffPer
	case Mystic:
		return MystEffPer
	case Sonic:
		return SonicEffPer
	}
	return -1
}
// Defense type has no equivalent