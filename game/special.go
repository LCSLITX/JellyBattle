package game

import (
	"fmt"
)

// Check out more struct based enums: https://threedots.tech/post/safer-enums-in-go/

// SPECIAL TYPE
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

// SPECIAL NAME
type specialName uint8

const (
	// Weapon
	Handbag    specialName = iota // 0
	Blast                         // 1
	Laser                         // 2
	AirStrike                     // 3
	RandomDrop                    // 4
	Nuke                          // 5
	// Buffs
	Invisible  // 6
	DeBonus    // 7
	Shield     // 8
	Freeze     // 9
	JellyRage  // 10
	TripleJump // 11
	// Utility
	Teleport // 12
	// Trap
	Mine // 13
	// Health
	Health     // 14
	SuperJelly // 15
)

func (s specialName) String() string {
	return [...]string{"Handbag", "Blast", "Laser", "Air Strike", "Random Drop", "Nuke", "Invisible", "De-Bonus", "Shield", "Freeze", "Jelly Rage", "Triple Jump", "Teleport", "Mine", "Health", "Super Jelly"}[s]
}

func (s specialName) Description() string {
	return [...]string{
		"Collect to injure an opponent within one square",                     // 0
		"Activate a blast that injures opponents within three squares",        // 1
		"Injure opponents with a vertical and horizontal laser",               // 2
		"Need help? Call in an air strike",                                    // 3
		"Collect to drop a random object on an opponent",                      // 4
		"Injure all opponents with the ultimate bomb",                         // 5
		"Become invisible and gain a strategic advantage",                     // 6
		"Disarm your opponents' bonuses",                                      // 7
		"The ultimate protection",                                             // 8
		"Freeze an opponent in their current position",                        // 9
		"Cause double damage to opponents' health",                            // 10
		"Collect to move three spaces in any direction",                       // 11
		"Escape imminent danger by transporting yourself across the keyboard", // 12
		"Avoid mines to keep your health",                                     // 13
		"Collect to restore your health",                                      // 14
		"Collect to restore your health to 100%",                              // 15
	}[s]
}

func (s specialName) EnumIndex() int {
	return int(s)
}

var specials Specials = Specials{
	// Weapon (Pink)
	{ID: uint8(specialName(0).EnumIndex()), Type: specialType(0).String(), Name: specialName(0).String(), Multiplier: 1, Description: specialName(0).Description()},
	{ID: uint8(specialName(1).EnumIndex()), Type: specialType(0).String(), Name: specialName(1).String(), Multiplier: 1, Description: specialName(1).Description()},
	{ID: uint8(specialName(2).EnumIndex()), Type: specialType(0).String(), Name: specialName(2).String(), Multiplier: 1, Description: specialName(2).Description()},
	{ID: uint8(specialName(3).EnumIndex()), Type: specialType(0).String(), Name: specialName(3).String(), Multiplier: 1, Description: specialName(3).Description()},
	{ID: uint8(specialName(4).EnumIndex()), Type: specialType(0).String(), Name: specialName(4).String(), Multiplier: 1, Description: specialName(4).Description()}, // TODO: The damage of this special varies 25, 35 and 50.
	{ID: uint8(specialName(5).EnumIndex()), Type: specialType(0).String(), Name: specialName(5).String(), Multiplier: 1, Description: specialName(5).Description()},
	// Buff (Blue)
	{ID: uint8(specialName(6).EnumIndex()), Type: specialType(1).String(), Name: specialName(6).String(), Multiplier: 1, Description: specialName(6).Description()},
	{ID: uint8(specialName(7).EnumIndex()), Type: specialType(1).String(), Name: specialName(7).String(), Multiplier: 1, Description: specialName(7).Description()},
	{ID: uint8(specialName(8).EnumIndex()), Type: specialType(1).String(), Name: specialName(8).String(), Multiplier: 1, Description: specialName(8).Description()},
	{ID: uint8(specialName(9).EnumIndex()), Type: specialType(1).String(), Name: specialName(9).String(), Multiplier: 1, Description: specialName(9).Description()},
	{ID: uint8(specialName(10).EnumIndex()), Type: specialType(1).String(), Name: specialName(10).String(), Multiplier: 1, Description: specialName(10).Description()},
	{ID: uint8(specialName(11).EnumIndex()), Type: specialType(1).String(), Name: specialName(11).String(), Multiplier: 1, Description: specialName(11).Description()},
	// Utility (Yellow)
	{ID: uint8(specialName(12).EnumIndex()), Type: specialType(2).String(), Name: specialName(12).String(), Multiplier: 1, Description: specialName(12).Description()},
	// Trap (Red)
	{ID: uint8(specialName(13).EnumIndex()), Type: specialType(3).String(), Name: specialName(13).String(), Multiplier: 1, Description: specialName(13).Description()},
	// Healing (Green)
	{ID: uint8(specialName(14).EnumIndex()), Type: specialType(4).String(), Name: specialName(14).String(), Multiplier: 1, Description: specialName(14).Description()},
	{ID: uint8(specialName(15).EnumIndex()), Type: specialType(4).String(), Name: specialName(15).String(), Multiplier: 1, Description: specialName(15).Description()},
}

// GenerateSpecial() returns a special
func GenerateSpecial() Special {
	rand := RandomNumber(ZERO_TO_FIFTEEN) // 16 possible random numbers
	s := specials[specialName(rand).EnumIndex()]
	if DebugModeSpecial() {
		fmt.Printf("%v: %+v\n\n", Trace(""), s)
	}
	return s
}
