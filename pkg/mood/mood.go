package mood

import (

)

type Mood int

const (
	D Mood = iota
	C
	B
	A
	S
	SS
	Unknown
)

var MoodString = [...]string {
	"d",
	"c",
	"b",
	"a",
	"s",
	"ss",
	"unknown",
}