package game

import (
	"fmt"

	"github.com/lucassauro/J.B.Remake/game/utils"
)

func NewPlayer(name string) IPLayer {
	player := &Player{
		Name: name,
		Rank: uint8(0),
	}

	if utils.DebugMode() {
		fmt.Printf("%v: %+v\n\n", utils.Trace(), player)
	}

	return player
}

func (player *Player) GetPlayer() Player {
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

// 	if utils.DebugMode() {
// 		fmt.Printf("%v: %+v\n\n", utils.Trace(), g)
// 	}

	// return g
// }
