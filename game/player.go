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

// func (player *Player) FindGroup(groups *Groups) Group {
// 	players := Players{
// 		{Name: player.Name, Rank: player.Rank},
// 	}
// 	// TODO: Implement the logic to find more players probably will use go routines with timeouts

// 	g := Group{
// 		ID:      "0",
// 		Players: players,
// 	}

// 	groups.Groups = append(groups.Groups, g)

// 	if DebugMode() {
// 		fmt.Printf("%v: %+v\n\n", Trace(), g)
// 	}

// return g
// }
