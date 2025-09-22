package character

import (
	"github.com/duhrice/bacsim/pkg/affiliation"
	"github.com/duhrice/bacsim/pkg/class"
	"github.com/duhrice/bacsim/pkg/role"
	"github.com/duhrice/bacsim/pkg/position"
)

type info struct {
	id Char
	name string
	school affiliation.School
	club affiliation.Club
	baseStarGrade int
	class class.ClassType
	role role.RoleType
	// summons 		<- implementation on hold for now
	position position.PositionType
	// to be continued...
}

var CharInfo = map[Char]info {
	
}