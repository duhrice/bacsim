package size

import (
	
)

func (s Size) String() string {
	return SizeString[s]
}

func StringtoS(s string) Size {
	for i, v := range SizeString {
		// Check if size is valid
		if v == s {
			// Found valid size
			return Size(i)
		}
	}
	// Invalid size found
	return None
}