package mood

import (

)

func (m Mood) String() string {
	return MoodString[m]
}

func StrToMood(s string) Mood {
	for i, v := range MoodString {
		// Check if mood is valid
		if v == s {
			// Found valid mood
			return Mood(i)
		}
	}
	// Invalid mood found
	return Unknown
}