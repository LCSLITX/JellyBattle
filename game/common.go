package game


// Instantiated classes.
var (
	GROUPS *Groups = NewGroups().GetGroups()
	PLAYERLIST PlayerList = NewPlayerList().GetPlayerList()
	GAMES Games = NewGames().GetGames()
)
