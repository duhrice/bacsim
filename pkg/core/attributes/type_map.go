package attributes

import "encoding/json"

// Map for Attack Type
type AttackTypeMap map[AttackType]float64
// Map for Defense Type
type DefenseTypeMap map[DefenseType]float64

// JSON Marshalling for attack type map
func (at AttackTypeMap) MarshalJSON() ([]byte, error) {
	// Map of the string representation of the attack type
	stringRep := make(map[string]float64)
	// Create the key value pairs 
	for key, value := range at {
		stringRep[AttackTypeString[key]] = value
	}

	return json.Marshal(stringRep)
}

// JSON Marshalling for attack type map
func (dt DefenseTypeMap) MarshalJSON() ([]byte, error) {
	// Map of the string representation of the attack type
	stringRep := make(map[string]float64)
	// Create the key value pairs 
	for key, value := range dt {
		stringRep[DefenseTypeString[key]] = value
	}

	return json.Marshal(stringRep)
}