package class

import (

)

type ClassType int

const (
	Striker ClassType = iota
	Special
	Unknown
)

var ClassTypeString = [...]string {
	"striker",
	"special",
	"unknown",
}
