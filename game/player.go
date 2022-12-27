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
		if v.Special == specialName.String(11) {
			player.JumpDistance = 3
		}
	}
	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(), player.JumpDistance)
	}
}

// TODO: Think about how its gonna work. Return a matrice maybe, or just return 4 limits.
func (player *Player) GetJumpArea() {
	// jd := player.JumpDistance
	// row, column := player.Position.Row, player.Position.Column

	// limit1 := row + jd
	// limit2 := row - jd
	// limit3 := column + jd
	// limit4 := column - jd
}

func (player *Player) Jump(p Position) {
	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(), *player)
	}

	player.Position.Row, player.Position.Column = p.Row, p.Column

}

