package game

// _____ BOARDS _____


// Players can jump 2 tiles to any direction. unless he has triple jump then its 3 tiles.
// Players start at row index 1.
// Special start at row index 3.
// Preview row is row index 5.
type IBoard interface {
	// GetBoard returns the board
	GetBoard() Board
	// GenerateBoard creates a board with specified number of rows and columns
	GenerateBoard(rows, columns uint8)

	// both useless after refactoring
	// RandomizeBoard populates a board with buttons
	// RandomizeBoard()
	// UpdateBoard updates the board buttons
	// UpdateBoard()

	// RoundRows rotates the board rows every round.
	RoundRows()
	// GeneratePreviewRow generate one row at once
	GeneratePreviewRow()

	// CountButtons return the quantity of buttons in the grid.
	countButtons() uint64
	// CountEmptyButtons return the quantity of empty and Fulfilled buttons.
	countEmptyButtons() uint64
}

type IButton interface {
	// RandomizeButton generate an empty button or special button
	RandomizeButton()
	// UpdateButton updates row and columns inside each button
	UpdateButton(row, column uint8)
	// EraseButtonSpecial updates a button by removing its special
	EraseButtonSpecial()
}



// _____ PLAYERS _____

// TODO: Implementation of players package. Need to decide Player and Dummy structures.
type IPLayer interface {
	// GetPlayer return a player.
	GetPlayer() Player
	GetJumpDistance()
	GetJumpArea()
	Jump(Position)
	// UseButton()
	// TakeDamage()
	// DoDamage()
	// ReplenishLife()
}

type IDummy interface{}

type IPlayerList interface {
	GetPlayerList() PlayerList
	AddPlayer(Player)
	RemovePlayer(Player)
	FindPlayer(p Player) (i int)
	GroupFourPlayers(*Groups) (Group, error)
}

// _____ GAMES _____

type IGame interface {
	// GetGames(Games) Game // Not sure yet
	StartGame() bool
	FinishGame() bool
	GetWinners() Group // Could also return Players []Player. Will be ordered by index. 0 is winner, 3 is last
}

// _____ GROUPS _____

type IGroups interface {
	GetGroups() Groups
	AddGroup(Group)
}
