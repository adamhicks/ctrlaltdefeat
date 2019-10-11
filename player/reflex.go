package player

import (
	"github.com/luno/reflex"
)

// Event is a reflex event type.
type Event = reflex.Event

// EventType is an int custom type.
type EventType int

// ReflexType returns the int representation of an EventType.
func (t EventType) ReflexType() int {
	return int(t)
}

// Ticketing event types are listed here.
const (
	EventTypeUnknown               EventType = 0
	EventTypePlayerRoundJoined     EventType = 1
	EventTypePlayerRoundExcluded   EventType = 2
	EventTypePlayerRoundJoining    EventType = 3
	EventTypePlayerRoundCollecting EventType = 4
	EventTypePlayerRoundCollected  EventType = 5
	EventTypePlayerRoundSubmitting EventType = 6
	EventTypePlayerRoundEnded      EventType = 7
	eventTypeSentinel              EventType = 8
)
