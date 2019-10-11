package player

import "time"

type RoundInfo struct {
	Rank  int64
	Parts []PartInfo
}

type PartInfo struct {
	Player string
	Part   int64
}
<<<<<<< HEAD

type PlayerRound struct {
	ID        int64
	RoundID   int64
	Status    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PlayerRoundStatus int

func (prs PlayerRoundStatus) Enum() int {
	return int(prs)
}

func (prs PlayerRoundStatus) ReflexType() int {
	return prs.Enum()
}

func (PlayerRoundStatus) ShiftStatus() {
}

var (
	PlayerRoundStatusUnknown         PlayerRoundStatus = 0
	PlayerRoundStatusRoundJoined     PlayerRoundStatus = 2
	PlayerRoundStatusRoundExcluded   PlayerRoundStatus = 3
	PlayerRoundStatusRoundJoining    PlayerRoundStatus = 1
	PlayerRoundStatusRoundCollecting PlayerRoundStatus = 4
	PlayerRoundStatusRoundCollected  PlayerRoundStatus = 5
	PlayerRoundStatusRoundSubmitting PlayerRoundStatus = 6
	PlayerRoundStatusRoundSubmitted  PlayerRoundStatus = 7
	PlayerRoundStatusRoundEnded      PlayerRoundStatus = 8
	PlayerRoundStatusSentinel        PlayerRoundStatus = 9
)

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
