package game

// Instantiate classes.
var (
	GROUPS               *Groups              = NewGroups().GetGroups()
	AVAILABLEPLAYERSLIST AvailablePlayersList = NewPlayerList().GetPlayerList()
	GAMES                Games                = NewGames().GetGames()
)
