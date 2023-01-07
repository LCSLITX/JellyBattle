package game

// Instantiate classes.
var (
	GROUPS     *Groups              = NewGroups().GetGroups()
	PLAYERLIST AvailablePlayersList = NewPlayerList().GetPlayerList()
	GAMES      Games                = NewGames().GetGames()
)
