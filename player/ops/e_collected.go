package ops

//SubmitRoundsForever
//Get PRs with state RoundCollected, get the round parts, if it's first SubmitRound and transition to RoundSubmitted
//ConsumeRoundSubmitsForever
//Listen for PRRoundSubmitted, get the round parts, if it's next SubmitRound and transition to RoundSubmitted
