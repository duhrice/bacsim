package attributes

import "encoding/json"

// Map for Attack Type
type AttackTypeMap map[AttackType]float64

// JSON Marchalling for student attack type map
func (at AttackTypeMap) MarshalJSON() ([]byte, error) {
	// Map of the string representation of the student attack type
	stringRep := make(map[string]float64)
	// Create the key value pairs 
	for key, value := range at {
		stringRep[AttackTypeString[key]] = value
	}

	return json.Marshal(stringRep)
}