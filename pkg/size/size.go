package size

import (

)

type Size int

const (
	Small Size = iota
	Medium
	Large
	None
)

var SizeString = [...]string {
	"small",
	"medium",
	"large",
	"n/a",
}