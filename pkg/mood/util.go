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

// Terrain combat mood multiplier
func MoodtoFloat(m Mood) float32 {
	switch m {
	case D:
		return 0.8
	case C:
		return 0.9
	case B:
		return 1.0
	case A:
		return 1.1
	case S:
		return 1.2
	case SS:
		return 1.3
	}
	// If it reaches this point, something is wrong
	return -1
}