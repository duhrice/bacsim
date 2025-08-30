package affiliation

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

func (c Club) String() string {
	return ClubString[c]
}

func StrToClub(s string) Club {
	for i, v := range ClubString {
		// Check if club is valid
		if v == s {
			// Found valid club
			return Club(i)
		}
	}
	// Invalid club found
	return None
}