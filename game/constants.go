package game

import "time"

const (
	DEFAULT_PLAYERS_NUMBER uint8 = 4
	DEFAULT_ROWS           uint8 = 5
	DEFAULT_COLUMNS        uint8 = 15

	DEFAULT_SPECIAL_FULFILLMENT uint8 = 70

	DEFAULT_INTERVAL = time.Second * 15 // doesn't make sense here

	DEFAULT_ROUND_TIME = 5  // in seconds
	DEFAULT_GAME_TIME  = 10 // in minutes

	ZERO_TO_NINETY_NINE int = 99
	ZERO_TO_FOUR        int = 4
	ZERO_TO_FIFTEEN     int = 15
)
