package defense

import (

)

func (dt DefenseType) String() string {
	return DefenseTypeString[dt]
}

func StringtoDT(s string) DefenseType {
	for i, v := range DefenseTypeString {
		// Check if defense type is valid
		if v == s {
			// Found valid defense type
			return DefenseType(i)
		}
	}
	// Invalid defense type found
	return Unknown
}