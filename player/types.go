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

type PlayerRound struct {
	ID        int64
	Status    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PlayerRoundStatus int

func (prs PlayerRoundStatus) Enum() int {
	return int(prs)
}

func (aus PlayerRoundStatus) ReflexType() int {
	return attachmentUploadReflex[aus].ReflexType()
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
)

// attachmentUploadReflex maps AttachmentUploadStatus types to reflex EventTypes.
var attachmentUploadReflex = map[PlayerRoundStatus]EventType{
	PlayerRoundStatusRoundJoined:     EventTypePlayerRoundJoined,
	PlayerRoundStatusRoundExcluded:   EventTypePlayerRoundExcluded,
	PlayerRoundStatusRoundJoining:    EventTypePlayerRoundJoining,
	PlayerRoundStatusRoundCollecting: EventTypePlayerRoundCollecting,
	PlayerRoundStatusRoundCollected:  EventTypePlayerRoundCollected,
	PlayerRoundStatusRoundSubmitting: EventTypePlayerRoundSubmitting,
	PlayerRoundStatusRoundEnded:      EventTypePlayerRoundEnded,
}
