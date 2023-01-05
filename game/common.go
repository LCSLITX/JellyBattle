package game

// Instantiate classes.
var (
	GROUPS *Groups = NewGroups().GetGroups()
	PLAYERLIST PlayerList = NewPlayerList().GetPlayerList()
	GAMES Games = NewGames().GetGames()
)
