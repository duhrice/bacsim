package equipment

import (

)

func (et EquipmentType) String() string {
	return EquipmentTypeString[et]
}

func StringtoET(s string) EquipmentType {
	for i, v := range EquipmentTypeString {
		// Check if equipment type is valid
		if v == s {
			// Found valid equipment type
			return EquipmentType(i)
		}
	}
	// Invalid equipment type found
	return Unknown
}