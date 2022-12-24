package boards

// _____ BOARD STRUCTS _____

// Struct button refers to each button in the the board grid.
type Button struct {
	Row       uint8 // Not sure if this will be used. I think yes because a button apparently must have a row.
	Column    uint8 // Not sure if this will be used. I think no because buttons will not change columns.
	Fulfilled bool  // When true button does contain a special. When false means Empty, or no special.
	Special   Special
}

// Struct Buttons refers to a specific ROW in the board grid.
type Row []Button

// Struct Special refers to spells that goes in each button that is not empty.
// TODO: Think about adding Charges attributes as some buffs have fixed charges values. Each charge lasts for one round.
type Special struct {
	ID         uint8
	Type       string // Weapon (Pink), Buff (Blue), Utility (Yellow), Trap (Red), Healing (Green) or Empty(No color).
	Name       string
	Multiplier uint8  // used only for healing and Trap. Other type of special have fixed value.
	// Effect     string // used only for buffs spells. Otherwise "".
	// Power      uint8  // used only for attack spells. Otherwise 0.
}

type Specials []Special

// Struct Board refers to the board grid.
type Board struct {
	RowNumber          uint8  // Quantity of rows.
	ColumnNumber       uint8  // Quantity of columns.
	NumberOfButtons    uint64 // Quantity of buttons in the grid.
	Players            uint8  // Quantity of players.
	SpecialFulfillment uint8  // Percentage of Empty spaces.
	Rows               []Row  // Bi-dimensional array of Buttons. In a grid of 5x20, ButtonRows attribute would be equivalent to [5][20]Buttons.
}
