package role

type RoleType int

const (
	Tank RoleType = iota
	Dealer
	Healer
	Support
	TacSupport		// Tactical support (T.S. for short)
	Unknown
)

var RoleTypeString = [...]string {
	"tank",
	"dealer",
	"healer",
	"support",
	"tac_support",	// Tactical support (T.S. for short)
	"unknown",
}