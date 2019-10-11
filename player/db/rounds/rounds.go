package rounds

import (
	"context"
	"database/sql"

	"github.com/luno/reflex/rsql"

	"github.com/adamhicks/ctrlaltdefeat/player"
)

// Call FSM functions here.

func Joining(ctx context.Context, dbc *sql.DB, roundID int64) (int64, error) {
	return fsm.Insert(ctx, dbc, joining{RoundID: roundID})
}

func Excluded(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundJoining,
		player.PlayerRoundStatusRoundExcluded, excluded{ID: id})
}

func Joined(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundJoining,
		player.PlayerRoundStatusRoundJoined, joined{ID: id})
}

func Collecting(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundJoined,
		player.PlayerRoundStatusRoundCollecting, joined{ID: id})
}

func Collected(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundCollecting,
		player.PlayerRoundStatusRoundCollected, joined{ID: id})
}

func Submitting(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundCollected,
		player.PlayerRoundStatusRoundSubmitting, joined{ID: id})
}

func Submitted(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundSubmitting,
		player.PlayerRoundStatusRoundSubmitted, joined{ID: id})
}

func EndedJoined(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundSubmitted,
		player.PlayerRoundStatusRoundEnded, joined{ID: id})
}

func EndedExcluded(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundExcluded,
		player.PlayerRoundStatusRoundEnded, joined{ID: id})
}
