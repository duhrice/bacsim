package position

import (

)

func (pt PositionType) String() string {
	return PositionTypeString[pt]
}

func StringtoPT(s string) PositionType {
	for i, v := range PositionTypeString {
		// Check if position type is valid
		if v == s {
			// Found valid position type
			return PositionType(i)
		}
	}
	// Invalid position type found
	return Unknown
}