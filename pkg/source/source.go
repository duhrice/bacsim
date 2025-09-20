package source

import (

)

type SourceType int

const (
	EX SourceType = iota
	Basic
	Enhanced
	Sub
	Unknown
)

var SourceTypeString = [...]string {
	"ex",
	"basic",
	"enhanced",
	"sub",
	"unknown",
}