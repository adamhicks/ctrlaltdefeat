package rounds

import (
	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/corverroos/unsure/engine/db/events"
	"github.com/luno/shift"
)

//go:generate shiftgen -inserter=joining -updaters=joined,excluded,collecting,collected,submitting,submitted  -table=player_rounds

var fsm = shift.NewFSM(events.GetTable()).
	Insert(player.PlayerRoundStatusRoundJoining, joining{},
		player.PlayerRoundStatusRoundJoined, player.PlayerRoundStatusRoundExcluded).
	Update(player.PlayerRoundStatusRoundExcluded, excluded{}, player.PlayerRoundStatusRoundEnded).
	Update(player.PlayerRoundStatusRoundJoined, joined{}, player.PlayerRoundStatusRoundCollecting).
	Update(player.PlayerRoundStatusRoundCollecting, collecting{}, player.PlayerRoundStatusRoundCollected.
	Update(player.PlayerRoundStatusRoundCollected, collected{}, player.PlayerRoundStatusRoundSubmitting).
	Update(player.PlayerRoundStatusRoundSubmitting, submitting{}, player.PlayerRoundStatusRoundSubmitted).
	Update(player.PlayerRoundStatusRoundSubmitted, submitted{}, player.PlayerRoundStatusRoundEnded).
	Build()

type joining struct {
}

type joined struct {
	ID int64
}

type excluded struct {
	ID int64
}

type collecting struct {
	ID int64
}

type collected struct {
	ID int64
}

type submitting struct {
	ID int64
}

type submitted struct {
	ID int64
}
