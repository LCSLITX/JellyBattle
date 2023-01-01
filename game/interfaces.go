package game

// _____ BOARDS _____

// Players can jump 2 tiles to any direction. unless he has triple jump then its 3 tiles.
// Players start at row index 1.
// Special start at row index 3.
// Preview row is row index 5.
type IBoard interface {
	countButtons() uint64	// CountButtons return the quantity of buttons in the grid.
	countEmptyButtons() uint64	// CountEmptyButtons return the quantity of empty and Fulfilled buttons.
	GenerateBoard(rows, columns uint8)	// GenerateBoard creates a board with specified number of rows and columns
	GeneratePreviewRow() // GeneratePreviewRow generate one row at once
	GetStartPosition()
	GetBoard() Board	// GetBoard returns the board
	RoundRows()	// RoundRows rotates the board rows every round.
}

type IButton interface {
	EraseButtonSpecial()	// EraseButtonSpecial updates a button by removing its special
	RandomizeButton()	// RandomizeButton generate an empty button or special button
	UpdateButton(row, column uint8)	// UpdateButton updates row and columns inside each button
}



// _____ PLAYERS _____

// TODO: Implementation of players package. Need to decide Player and Dummy structures.
type IPLayer interface {
	// DoDamage()
	GetPlayer() Player
	GetJumpDistance()
	GetJumpArea()
	JumpTo()
	NextJump(Position)
	// ReplenishLife()
	// TakeDamage()
	// UseButton()
}

type IDummy interface{}

type IPlayerList interface {
	AddPlayer(Player)
	FindPlayer(p Player) (i int)
	GetPlayerList() PlayerList
	GroupFourPlayers(*Groups) (Group, error)
	RemovePlayer(Player)
}

// _____ GAMES _____

type IGame interface {
	FinishGame() bool
	// GetGames(Games) Game // Not sure yet
	GetGame() Game
	GetWinners() Group // Could also return Players []Player. Will be ordered by index. 0 is winner, 3 is last
	StartGame() bool
}

type IGames interface {
	AddGame(Game)
	GetGames() Games
}

// _____ GROUPS _____

type IGroups interface {
	AddGroup(Group)
	GetGroups() *Groups
	GetFirstGroup() Group
}
