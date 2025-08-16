// Student attack types
package attributes

import (
	"encoding/json"
	"errors"
	"strings"
)

// String representing student attack type
type AttackType int

const (
	Explosive AttackType = iota
	Piercing
	Mystic
	Sonic
	Normal
	Unknown
)

// Allow attack types to be printed as strings
func (at AttackType) String() string {
	return AttackTypeString[at]
}

// JSON marshalling for student attack type
func (at AttackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttackTypeString[at])
}

// JSON unmarshalling for student attack type
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

// Student attack type as strings
var AttackTypeString = [...]string {
	"explosive",
	"piercing",
	"mystic",
	"sonic",
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

// One more function here for stat increases