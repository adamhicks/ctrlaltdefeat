package ops

//StartMatchesForever
//Check for number of rounds not in RoundEnded state, if == 0, try to start a match

//ConsumeMatchEventsForever
//Listen for MatchEnded event, try to start a match

//StartRoundsForever
//Listen for EventTypeRoundJoin event and create a PlayerRound(PR) object
