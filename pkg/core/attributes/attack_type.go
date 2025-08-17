package attributes

import (
	"encoding/json"
	"errors"
	"strings"
)

// Int representing student and enemy attack type
type AttackType int

// All student and enemy attack types
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

// Allow attack types to be printed as strings
func (at AttackType) String() string {
	return AttackTypeString[at]
}

// JSON marshalling for attack type
func (at AttackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttackTypeString[at])
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
	return errors.New("unrecognized student attack type")
}

// Attack types as strings
var AttackTypeString = [...]string {
	"explosive",
	"piercing",
	"mystic",
	"sonic",
	"siege",
	"normal",
	"unknown",
}

// Convert a attack type string to attack type
func StringtoAT(s string) AttackType {
	for i, v := range AttackTypeString {
		// Check if attack type is valid
		if v == s {
			// Found valid attack type
			return AttackType(i)
		}
	}
	// Invalid attack type found
	return Unknown
}

// Convert attack type to attack type effectiveness
func AtkTypeToEff(at AttackType) Stat {
	switch at {
	case Explosive:
		return ExplosiveEff
	case Piercing:
		return PiercingEff
	case Mystic:
		return MysticEff
	case Sonic:
		return SonicEff
	}
	return -1
}