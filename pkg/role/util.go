package role

import (
	
)

func (rt RoleType) String() string {
	return RoleTypeString[rt]
}

func StringtoRT(s string) RoleType {
	for i, v := range RoleTypeString {
		// Check if role type is valid
		if v == s {
			// Found valid role type
			return RoleType(i)
		}
	}
	// Invalid role type found
	return Unknown
}