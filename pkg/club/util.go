package club

import (

)

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