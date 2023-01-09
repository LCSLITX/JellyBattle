# JellyBattle
Remake of a classical flash game called Jelly Battle.


Priorities:
  -[X] Distribute players in each start position
  -[X] Debug groups and find out why getGroups is returning an empty array while addGroup is apparently working.
  -[X] Players can move in a game.
  -[ ] Think about how CauseDamage is gonna work

  -[ ] Continue debuging to check why panics
  ```
  [playersList.go:21 - game.(*PlayerList).GetPlayerList]: [{Name:P4 Rank:0 Life:0 Experience:0 GamesPlayed:0 JumpDistance:0 Position:{Row:0 Column:0} JumpPosition:{Row:0 Column:0} Buffs:[]} {Name:P2 Rank:0 Life:0 Experience:0 GamesPlayed:0 JumpDistance:0 Position:{Row:0 Column:0} JumpPosition:{Row:0 Column:0} Buffs:[]} {Name:P3 Rank:0 Life:0 Experience:0 GamesPlayed:0 JumpDistance:0 Position:{Row:0 Column:0} JumpPosition:{Row:0 Column:0} Buffs:[]}]
  panic: runtime error: index out of range [-1]
  ```

  -[ ] Use websocket when a game is started. It should transmit information to all players in that.
  -[ ] Implement in-game chat.
  -[ ] Player can move only within its JumpDistance. Maybe it will be a proper limitation of front-end.
  -[ ] Game runs from start to end, with rounds counting up and players moving in each turn.
  -[ ] Remove commentaries in main.go file. Those validations should've made with unit or integration tests.

  -[ ] Think about using websocket for searching for a game. Player connects to a list of available players and gets removed if it disconnects. It is a possible better approach than the current one.

  -[ ] Optimise structs