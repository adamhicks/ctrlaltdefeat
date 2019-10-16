package rounds

import (
	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/adamhicks/ctrlaltdefeat/player/events"
	"github.com/luno/shift"
)

//go:generate shiftgen -inserter=joining -updaters=updateRound  -table=player_rounds

var fsm = shift.NewFSM(events.GetTable()).
	Insert(player.PlayerRoundStatusRoundJoining, joining{},
		player.PlayerRoundStatusRoundJoined, player.PlayerRoundStatusRoundExcluded).
	Update(player.PlayerRoundStatusRoundExcluded, updateRound{}, player.PlayerRoundStatusRoundEnded).
	Update(player.PlayerRoundStatusRoundJoined, updateRound{}, player.PlayerRoundStatusRoundCollecting).
	Update(player.PlayerRoundStatusRoundCollecting, updateRound{}, player.PlayerRoundStatusRoundCollected).
	Update(player.PlayerRoundStatusRoundCollected, updateRound{}, player.PlayerRoundStatusRoundSubmitting).
	Update(player.PlayerRoundStatusRoundSubmitting, updateRound{}, player.PlayerRoundStatusRoundSubmitted).
	Update(player.PlayerRoundStatusRoundSubmitted, updateRound{}, player.PlayerRoundStatusRoundEnded).
	Update(player.PlayerRoundStatusRoundEnded, updateRound{}).
	Build()

type joining struct {
	RoundID int64
}

type updateRound struct {
	ID int64
}
