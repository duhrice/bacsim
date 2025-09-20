package source

import (

)

func (st SourceType) String() string {
	return SourceTypeString[st]
}

func StringtoST(s string) SourceType {
	for i, v := range SourceTypeString {
		// Check if source type is valid
		if v == s {
			// Found source role type
			return SourceType(i)
		}
	}
	// Invalid source type found
	return Unknown
}