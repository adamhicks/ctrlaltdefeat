package player

type RoundInfo struct {
	Rank  int64
	Parts []PartInfo
}

type PartInfo struct {
	Player string
	Part   int64
}

type RoundParts struct {
	ID       int64
	MatchID  int64
	RoundID  int64
	PlayerID string
	Rank     int64
	P1Part   int64
	P2Part   int64
	P3Part   int64
	P4Part   int64
}
