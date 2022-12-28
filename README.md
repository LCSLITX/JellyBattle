# J.B.Remake
Remake of a classical flash game called Jelly Battle.



Priorities:
  -[X] Distribute players in each start position
  -[X] Debug groups and find out why getGroups is returning an empty array while addGroup is apparently working.
  -[X] Players can move in a game.
  -[] Continue debbuging to check why panics
  ```
  [playersList.go:21 - game.(*PlayerList).GetPlayerList]: [{Name:P4 Rank:0 Life:0 Experience:0 GamesPlayed:0 JumpDistance:0 Position:{Row:0 Column:0} JumpPosition:{Row:0 Column:0} Buffs:[]} {Name:P2 Rank:0 Life:0 Experience:0 GamesPlayed:0 JumpDistance:0 Position:{Row:0 Column:0} JumpPosition:{Row:0 Column:0} Buffs:[]} {Name:P3 Rank:0 Life:0 Experience:0 GamesPlayed:0 JumpDistance:0 Position:{Row:0 Column:0} JumpPosition:{Row:0 Column:0} Buffs:[]}]
  panic: runtime error: index out of range [-1]
  ```
  -[] Code cmd.go. It is supposed to be a terminal application to emulate the game and turn easier the process of watching it.
  -[] Player can move only within its JumpDistance
  -[] Game runs from start to end, with rounds counting up and players moving in each turn.
