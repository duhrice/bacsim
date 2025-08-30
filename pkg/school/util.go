package school

import (

)

func (s School) String() string {
	return SchoolString[s]
}

func StrToSchool(s string) School {
	for i, v := range SchoolString {
		// Check if school is valid
		if v == s {
			// Found valid school
			return School(i)
		}
	}
	// Invalid school found
	return Unknown
}