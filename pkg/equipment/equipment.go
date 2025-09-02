package equipment

import (

)

type EquipmentType int

const (
	Hat EquipmentType = iota
    Gloves
    Shoes
    Bag
    Badge
    Hairpin
    Amulet
    Necklace
    Wristwatch
    UnqiueItem
	Unknown
)

var EquipmentTypeString = [...]string {
	"hat",
    "gloves",
    "shoes",
    "bag",
    "badge",
    "hairpin",
    "amulet",
    "necklace",
    "wristwatch",
    "unqiue_item",
	"unknown",
}