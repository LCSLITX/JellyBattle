package game

import (
	"fmt"
	"reflect"
)

// NewPlayerList() constructor returns a PlayerList instance.
func NewPlayerList() IAvailablePlayersList {
	p := &AvailablePlayersList{}

	if DebugModePlayerList() {
		fmt.Printf("%v: %+v\n\n", Trace(), p)
	}

	return p
}

// GetPlayerList() returns the playerList.
func (playerList *AvailablePlayersList) GetPlayerList() AvailablePlayersList {
	if DebugModePlayerList() {
		fmt.Printf("%v: %+v\n\n", Trace(), *playerList)
	}

	return *playerList
}

// AddPlayer() adds a player into the playerList.
func (playerList *AvailablePlayersList) AddPlayer(p Player) {
	// In case of doubt about next line: https://stackoverflow.com/questions/74915781/go-appending-elements-to-slice-of-struct
	*playerList = append(*playerList, p)

	if DebugModePlayerList() {
		fmt.Printf("%v: %+v\n\n", Trace(), *playerList)
	}
}

// RemovePlayer() removes the given player from the playerList by overriding it with the last player (index -1) and removing the last index.
func (playerList *AvailablePlayersList) RemovePlayer(p Player) {
	if len(playerList.GetPlayerList()) == 0 {
		return
	}

	i := playerList.FindPlayer(p)
	// In case of doubt about next lines:
	// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
	// https://stackoverflow.com/questions/38013922/slicing-a-slice-pointer-passed-as-argument
	(*playerList)[i] = (*playerList)[len(*playerList)-1]
	(*playerList) = (*playerList)[:len(*playerList)-1]

	if DebugModePlayerList() {
		fmt.Printf("%v: %+v\n\n", Trace(), *playerList)
	}
}

// FindPlayer() returns the index of a given player in the PlayerList. If not present returns -1.
func (playerList *AvailablePlayersList) FindPlayer(p Player) (i int) {
	i = -1
	for k, v := range *playerList {
		if reflect.DeepEqual(v, p) { // attention: it may have a bug if a player is passed with different buff. So mind where you take players from.
			i = k
			break
		}
	}
	return
}

// GetFourPlayers() get the first 4 players in the list and form a group with them.
func (playerList *AvailablePlayersList) GroupFourPlayers(groups *Groups) (Group, error) {
	if len(*playerList) >= 3 {
		p1, p2, p3, p4 := (*playerList)[0], (*playerList)[1], (*playerList)[2], (*playerList)[3]
		players := Players{
			p1,
			p2,
			p3,
			p4,
		}

		g := Group{
			ID:      GenerateID(), // TODO: Function to generate ID.
			Players: players,
		}

		// TODO: Implement an assynchronous function to verify if ID already exists and change it if its the case.

		// TODO:  Think if its necessary to Implement MUTEX.
		go func() {
			for _, v := range g.Players {
				playerList.RemovePlayer(v)
			}
		}()

		groups.AddGroup(g)

		return g, nil
	}

	return Group{}, fmt.Errorf("Not enough players to play.")
}
