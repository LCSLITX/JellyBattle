package game

import (
	"fmt"

	"github.com/lucassauro/J.B.Remake/game/utils"
)


func NewPlayerList() IPlayerList {
	p := &PlayerList{}

	if utils.DebugMode() {
		fmt.Printf("%v: %+v\n\n", utils.Trace(), p)
	}

	return p
}


// GetPlayerList returns the playerList.
func (playerList *PlayerList) GetPlayerList() PlayerList {
	if utils.DebugMode() {
		fmt.Printf("%v: %+v\n\n", utils.Trace(), *playerList)
	}

	return *playerList
}


// AddPlayer adds a player into the playerList.
// https://stackoverflow.com/questions/74915781/go-appending-elements-to-slice-of-struct
func (playerList *PlayerList) AddPlayer(p Player) {
	*playerList = append(*playerList, p)

	if utils.DebugMode() {
		fmt.Printf("%v: %+v\n\n", utils.Trace(), *playerList)
	}
}


// RemovePlayer removes the given player from the playerList by overriding it with the last player (index -1) and removing the last index.
// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
// https://stackoverflow.com/questions/38013922/slicing-a-slice-pointer-passed-as-argument
func (playerList *PlayerList) RemovePlayer(p Player) {
	i := playerList.FindPlayer(p)
	(*playerList)[i] = (*playerList)[len(*playerList) - 1]
	(*playerList) = (*playerList)[:len(*playerList) - 1]

	if utils.DebugMode() {
		fmt.Printf("%v: %+v\n\n", utils.Trace(), *playerList)
	}
}


// FindPlayer returns the index of a given player in the PlayerList. If not present returns -1.
func (playerList *PlayerList) FindPlayer(p Player) (i int) {
	i = -1
	for k, v := range *playerList {
		if v == p {
			i = k
			break
		}
	}
	return
}

// GetFourPlayers get the first 4 players in the list and form a group with them.
func (playerList *PlayerList) GetFourPlayers() (Group, error) {
	if len(*playerList) >= 3 {
		p1,	p2, p3, p4 := (*playerList)[0], (*playerList)[1], (*playerList)[2], (*playerList)[3]
		players := Players{
			p1,
			p2,
			p3,
			p4,
		}
		g := Group{
			ID: "0", // TODO: Function to generate ID.
			Players: players,
		}
		return g, nil
	}
	return Group{}, fmt.Errorf("Not enough players to play.")
}