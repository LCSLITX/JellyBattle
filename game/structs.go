package game

// _____ BOARDS _____

// Struct button refers to each button in the the board grid.
type Button struct {
	Row       uint8 // Not sure if this will be used. I think yes because a button apparently must have a row.
	Column    uint8 // Not sure if this will be used. I think no because buttons will not change columns.
	Fulfilled bool  // When true button does contain a special. When false means Empty, or no special.
	Special   Special
}

// Struct Row refers to a specific button row in the board grid.
type Row []Button

// Struct Special refers to powers that goes in each button that is not empty.
type Special struct {
	ID         uint8
	Type       string // Weapon (Pink), Buff (Blue), Utility (Yellow), Trap (Red), Healing (Green) or Empty(No color).
	Name       string
	Multiplier uint8 // used only for healing. Other type of special have fixed value.
	// Charges    uint8 // TODO: Think about adding Charges attributes as some buffs have fixed charges values. Each charge lasts for one round.
	// Order    uint8 // TODO: Think about adding Order attributes. Each special has a priority over others. Health -> Buffs -> ... -> Weapon
	// TODO: If players try to use the same special (at same time) each got a number from 1 to 10 and the player with higher number get the special while other suffers damage and falls to the button next to it.
}

type Specials []Special

// Struct Board refers to the board grid.
type Board struct {
	RowNumber          uint8  // Quantity of rows.
	ColumnNumber       uint8  // Quantity of columns.
	NumberOfButtons    uint64 // Quantity of buttons in the grid.
	PlayersNumber      uint8  // Quantity of players.
	SpecialFulfillment uint8  // Percentage of Empty spaces.
	Rows               []Row  // Bi-dimensional array of Buttons. In a grid of 5x20, ButtonRows attribute would be equivalent to [5][20]Buttons.
}

// _____ PLAYERS _____

// TODO: Decide future of Dummies
type Dummy struct {
	Name string
	Rank uint8
}

type Player struct {
	Name string
	Rank uint8
}

type Players []Player // Supposed to be used for a specific game, players will leave PlayerList to form a Players group to join game.

type PlayerList []Player // Supposed to have all the players available to play a game. 


type Group struct {
	ID      string
	Players Players
}

type Groups []Group

// _____ GAMES _____

type Game struct {
	ID    string
	Board Board
	Group Group
}
