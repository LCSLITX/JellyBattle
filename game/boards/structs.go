package boards

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
