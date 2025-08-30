package school

import (

)

type School int

// No playable students from the following schools
// Kivotos, Kronos, Odyssey Maritime
const (
	Abydos School = iota	// Abydos High School
	Arius					// Arius Branch School
	Gehenna					// Gehenna Academy
	Highlander				// Highlander Railroad Academy
	Hyakkiyako				// Hyakkiyako Alliance Academy
	Kivotos					// For General Student Council and SCHALE
	Kronos					// Kronos School of Journalism
	Millennium				// Millennium Science School
	// OdysseyMaritime		// Odyssey Maritime School. Not yet expanded upon in-game yet. Leave commented out
	RedWinter				// Red Winter Federal Academy
	Shanhaijing				// Shanhaijing Senior Secondary School
	Tokiwadai				// Tokiwadai Middle School. For limited Railgun Collab
	Trinity					// Trinity General School
	Valkyrie				// Valkyrie Police Academy
	SRT						// SRT Special Academy
	WildHunt				// Wild Hunt Academy of Arts
	Etc						// For "Hatsune Miku"
	Unknown
)

var SchoolString = [...]string {
	"abydos",				// Abydos High School
	"arius",				// Arius Branch School
	"gehenna",				// Gehenna Academy
	"highlander",			// Highlander Railroad Academy
	"hyakkiyako",			// Hyakkiyako Alliance Academy
	"kivotos",				// For General Student Council and SCHALE
	"kronos",				// Kronos School of Journalism
	"millennium",			// Millennium Science School
	// "odyssey_maritime",	// Odyssey Maritime School. Not yet expanded upon in-game yet. Leave commented out
	"red_winter",			// Red Winter Federal Academy
	"shanhaijing",			// Shanhaijing Senior Secondary School
	"tokiwadai",			// Tokiwadai Middle School. For limited Railgun Collab
	"trinity",				// Trinity General School
	"valkyrie",				// Valkyrie Police Academy
	"srt",					// SRT Special Academy
	"wild_hunt",			// Wild Hunt Academy of Arts
	"etc",					// For "Hatsune Miku"
	"unknown",
}