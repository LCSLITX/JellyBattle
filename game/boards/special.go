package boards

import (
	"github.com/lucassauro/J.B.Remake/game/utils"
)

// Check out more struct based enums: https://threedots.tech/post/safer-enums-in-go/

type specialType uint8

const (
	W specialType = iota
	B
	U
	T
	H
)

func (s specialType) String() string {
	return [...]string{"Weapon", "Buff", "Utility", "Trap", "Healing"}[s]
}

func (s specialType) EnumIndex() int {
	return int(s)
}

type specialName uint8

const (
	Handbag specialName = iota
	Blast
	Laser
	AirStrike
	RandomDrop
	Nuke
	Invisible
	DeBonus
	Shield
	Freeze
	JellyRage
	TripleJump
	Teleport
	Trap
	Health
	SuperJelly
)

func (s specialName) String() string {
	return [...]string{"Handbag", "Blast", "Laser", "Air Strike", "Random Drop", "Nuke", "Invisible", "De-Bonus", "Shield", "Freeze", "Jelly Rage", "Triple Jump", "Teleport", "Trap", "Health", "Super Jelly"}[s]
}

func (s specialName) EnumIndex() int {
	return int(s)
}

var specials Specials = Specials{
	{ID: uint8(specialName(0).EnumIndex()), Type: specialType(0).String(), Name: specialName(0).String(), Multiplier: 1},
	{ID: uint8(specialName(1).EnumIndex()), Type: specialType(0).String(), Name: specialName(1).String(), Multiplier: 1},
	{ID: uint8(specialName(2).EnumIndex()), Type: specialType(0).String(), Name: specialName(2).String(), Multiplier: 1},
	{ID: uint8(specialName(3).EnumIndex()), Type: specialType(0).String(), Name: specialName(3).String(), Multiplier: 1},
	{ID: uint8(specialName(4).EnumIndex()), Type: specialType(0).String(), Name: specialName(4).String(), Multiplier: 1}, // TODO: The hit of this special varies 25, 35 and 50.
	{ID: uint8(specialName(5).EnumIndex()), Type: specialType(0).String(), Name: specialName(5).String(), Multiplier: 1},
	// Buff (Blue)
	{ID: uint8(specialName(6).EnumIndex()), Type: specialType(1).String(), Name: specialName(6).String(), Multiplier: 1},
	{ID: uint8(specialName(7).EnumIndex()), Type: specialType(1).String(), Name: specialName(7).String(), Multiplier: 1},
	{ID: uint8(specialName(8).EnumIndex()), Type: specialType(1).String(), Name: specialName(8).String(), Multiplier: 1},
	{ID: uint8(specialName(9).EnumIndex()), Type: specialType(1).String(), Name: specialName(9).String(), Multiplier: 1},
	{ID: uint8(specialName(10).EnumIndex()), Type: specialType(1).String(), Name: specialName(10).String(), Multiplier: 1},
	{ID: uint8(specialName(11).EnumIndex()), Type: specialType(1).String(), Name: specialName(11).String(), Multiplier: 1},
	// Utility (Yellow)
	{ID: uint8(specialName(12).EnumIndex()), Type: specialType(2).String(), Name: specialName(12).String(), Multiplier: 1},
	// Trap (Red)
	{ID: uint8(specialName(13).EnumIndex()), Type: specialType(3).String(), Name: specialName(13).String(), Multiplier: 1} ,
	// Healing (Green)
	{ID: uint8(specialName(14).EnumIndex()), Type: specialType(4).String(), Name: specialName(14).String(), Multiplier: 1},
	{ID: uint8(specialName(15).EnumIndex()), Type: specialType(4).String(), Name: specialName(15).String(), Multiplier: 1},
}

// GenerateSpell returns a spell
func GenerateSpecial() Special {
	rand := utils.RandomNumber(ZERO_TO_FIFTEEN) // 5 possible random numbers
	return specials[specialName(rand).EnumIndex()]
}
