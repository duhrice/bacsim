package modifier

import (

)

func (m ModType) String() string {
	return ModTypeString[m]
}

func StrToMod(s string) ModType {
	for i, v := range ModTypeString {
		// Check if mod type is valid
		if v == s {
			// Found valid mod type
			return ModType(i)
		}
	}
	// Invalid mod type found
	return -1
}