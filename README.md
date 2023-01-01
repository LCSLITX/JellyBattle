# J.B.Remake
Remake of a classical flash game called Jelly Battle.


Priorities:
  -[X] Distribute players in each start position
  -[X] Debug groups and find out why getGroups is returning an empty array while addGroup is apparently working.
  -[X] Players can move in a game.
  -[ ] Continue debuging to check why panics
  ```
  [playersList.go:21 - game.(*PlayerList).GetPlayerList]: [{Name:P4 Rank:0 Life:0 Experience:0 GamesPlayed:0 JumpDistance:0 Position:{Row:0 Column:0} JumpPosition:{Row:0 Column:0} Buffs:[]} {Name:P2 Rank:0 Life:0 Experience:0 GamesPlayed:0 JumpDistance:0 Position:{Row:0 Column:0} JumpPosition:{Row:0 Column:0} Buffs:[]} {Name:P3 Rank:0 Life:0 Experience:0 GamesPlayed:0 JumpDistance:0 Position:{Row:0 Column:0} JumpPosition:{Row:0 Column:0} Buffs:[]}]
  panic: runtime error: index out of range [-1]
  ```

  -[ ] Player can move only within its JumpDistance
  -[ ] Game runs from start to end, with rounds counting up and players moving in each turn.
  -[ ] Chat works.
  -[ ] Use websocket when a game is started. It should transmit information to all players in that.