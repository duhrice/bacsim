package position

type PositionType int

const (
	Front PositionType = iota
	Middle
	Back
	Unknown
)

var PositionTypeString = [...]string {
	"front",
	"middle",
	"back",
	"unknown",
}