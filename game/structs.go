package game

// How to optimise struct memory allocation in golang:
// https://medium.com/techverito/golang-struct-size-and-memory-optimisation-b46b124f008d
// https://stackoverflow.com/questions/2113751/sizeof-struct-in-go
// https://stackoverflow.com/questions/73211746/does-go-use-something-like-space-padding-for-structs

type Board struct { // Struct Board refers to the board grid.
	RowNumber          uint8  // Quantity of rows.
	ColumnNumber       uint8  // Quantity of columns.
	NumberOfButtons    uint64 // Quantity of buttons in the grid.
	PlayersNumber      uint8  // Quantity of players.
	SpecialFulfillment uint8  // Percentage of Empty spaces.
	PreviewRow         Row
	Rows               []Row // Bi-dimensional array of Buttons. In a grid of 5x20, ButtonRows attribute would be equivalent to [5][20]Buttons.
	StartPositions     Positions
}

type Button struct { // Struct button refers to each button in the the board grid.
	Row       uint8 // Not sure if this will be used. I think yes because a button apparently must have a row.
	Column    uint8 // Not sure if this will be used. I think no because buttons will not change columns.
	Fulfilled bool  // When true button does contain a special. When false means Empty, or no special.
	Special   Special
}

type Dummy struct { // TODO: Decide future of Dummies
	Name string
	Rank uint8
}

var Finish chan bool // outside Game because its not compatible with enconding/json.NewDecoder().Decode().

type Game struct { // Struct Game refers to a game, composed by a group of players and a board.
	ID       string
	Started  bool // Not sure yet
	Finished bool // Not sure yet
	Deaths   uint8
	Chat     []string // Not sure yet
	Board    Board
	Group    Group
	// Timer    time.Duration // Not sure yet
}

type Games map[string]Game

type Group struct {
	ID      string
	Players Players
}

type Groups []Group

type Player struct {
	ID           string
	Name         string
	Rank         uint8
	Life         uint8
	JumpDistance uint8
	Connected    bool // To be used on websocket connection
	GamesPlayed  uint
	Experience   uint64
	Position     Position
	JumpPosition Position
	Buffs        SpecialCharges
}


// TODO:  Think about the possibility of changing array for maps. Maybe it'll be better.
type Players map[string]*Player // Supposed to be used for a specific game, players will leave PlayerList to form a Players group to join game.

// TODO:  Think about the possibility of changing array for maps. Maybe it'll be better.
type AvailablePlayersList []*Player // Supposed to have all the players available to play a game.

type Position struct {
	Row    uint8 `json:"r"`    // x
	Column uint8 `json:"c"` // y
}

type Positions []Position

type Row []Button // Struct Row refers to a specific button row in the board grid.

type Special struct { // Struct Special refers to powers that goes in each button that is not empty.
	ID          uint8
	Type        string // Weapon (Pink), Buff (Blue), Utility (Yellow), Trap (Red), Healing (Green) or Empty(No color).
	Name        string
	Multiplier  uint8 // used only for healing. Other type of special have fixed value.
	Damage      uint8
	Description string
	Charges     uint8 // TODO: Think about adding Charges attributes as some buffs have fixed charges values. Each charge lasts for one round.
	// Order    uint8 // TODO: Think about adding Order attributes. Each special has a priority over others. Utility -> Health -> Buffs -> ... -> Weapon
	// TODO: If players try to use the same special (at same time) each got a number from 1 to 10 and the player with higher number get the special while other suffers damage and falls to the button next to it.
}

type SpecialCharge struct {
	SpecialName string
	Charges     uint8
}

type SpecialCharges []SpecialCharge

type Specials []Special
