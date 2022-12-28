package game

import (
	"fmt"
)

func NewPlayer(name string) IPLayer {
	player := &Player{
		Name: name,
		Rank: uint8(0),
	}

	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(), player)
	}

	return player
}

func (player *Player) GetPlayer() Player {
	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(), *player)
	}

	return *player
}

func (player *Player) GetJumpDistance()  {
	player.JumpDistance = 2
	for _, v := range player.Buffs {
		if v.SpecialName == specialName.String(11) {
			player.JumpDistance = 3
		}
	}
	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(), player.JumpDistance)
	}
}

// TODO: Think about how its gonna work. Return a matrice maybe, or just return 4 limits.
// The horizontal order of buttons (rows) is inverse to the cartesian plane.
func (player *Player) GetJumpArea() {
	// jd := player.JumpDistance
	// row, column := player.Position.Row, player.Position.Column

	// superiorRowLimit := row - jd
	// inferiorRowLimit := row + jd
	// leftColumnLimit := column - jd
	// rightColumnLimit := column + jd
}

// JumpTo Position is the function that changes player button.
func (player *Player) JumpTo() {
	player.Position.Row, player.Position.Column = player.JumpPosition.Row, player.JumpPosition.Column

	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(), *player)
	}
}

// NextJump Position is the function that receives the position of nextjump.
func (player *Player) NextJump(p Position) {
	player.JumpPosition.Row, player.JumpPosition.Column = p.Row, p.Column
	
	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(), *player)
	}
}

func (player *Player) TakeDamage(s Special) {
	if s.Type == "Weapon" {
		player.Life -= s.Damage
	} else if s.Type == "Healing" {
		player.Life += s.Damage
	}
}

func (player *Player) DoDamage(s Special, p Players) {

}