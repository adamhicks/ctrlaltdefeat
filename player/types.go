package player

type RoundInfo struct {
	Rank  int64
	Parts []PartInfo
}

type PartInfo struct {
	Player string
	Part   int64
}
