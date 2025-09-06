package character

func (c Char) String() string {
	return CharString[c]
}

func StrToChar(s string) Char {
	for i, v := range CharString {
		// Check if character is valid
		if v == s {
			// Found valid character
			return Char(i)
		}
	}
	// Invalid character found
	return -1
}