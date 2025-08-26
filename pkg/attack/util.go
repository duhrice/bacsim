package attack

import (

)

func (at AttackType) String() string {
	return AttackTypeString[at]
}

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