package weapon

import (

)

type WeaponType int

const (
	SG WeaponType = iota	// Shotgun
	SMG						// Submachine gun
	AR						// Assault rifle
	GL						// Grenade launcher
	HG						// Handgun
	SR						// Sniper rifle
	RG						// Railgun
	MG						// Machine gun
	RL						// Rocket launcher
	MT						// Mortar
	FT						//Flamethrower
	Unknown
)

var WeaponTypeString = [...]string {
	"sg",		// Shotgun
	"smg",		// Submachine gun
	"ar",		// Assault rifle
	"gl",		// Grenade launcher
	"hg",		// Handgun
	"sr",		// Sniper rifle
	"rg",		// Railgun
	"mg",		// Machine gun
	"rl",		// Rocket launcher
	"mt",		// Mortar
	"ft",		//Flamethrower
	"unknown",
}