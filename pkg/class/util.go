package class

import (

)

func (ct ClassType) String() string {
	return ClassTypeString[ct]
}

func StringtoCT(s string) ClassType {
	for i, v := range ClassTypeString {
		// Check if class type is valid
		if v == s {
			// Found valid class type
			return ClassType(i)
		}
	}
	// Invalid class type found
	return Unknown
}