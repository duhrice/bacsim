package weapon

import (

)

func (wt WeaponType) String() string {
	return WeaponTypeString[wt]
}

func StringtoWT(s string) WeaponType {
	for i, v := range WeaponTypeString {
		// Check if weapon type is valid
		if v == s {
			// Found valid weapon type
			return WeaponType(i)
		}
	}
	// Invalid weapon type found
	return Unknown
}