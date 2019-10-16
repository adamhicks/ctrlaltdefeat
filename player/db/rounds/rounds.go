package rounds

import (
	"context"
	"database/sql"

	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/luno/reflex/rsql"
)

// Call FSM functions here.

func Joining(ctx context.Context, dbc *sql.DB, roundID int64) (int64, error) {
	return fsm.Insert(ctx, dbc, joining{RoundID: roundID})
}

func Excluded(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundJoining,
		player.PlayerRoundStatusRoundExcluded, updateRound{ID: id})
}

func Joined(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundJoining,
		player.PlayerRoundStatusRoundJoined, updateRound{ID: id})
}

func Collecting(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundJoined,
		player.PlayerRoundStatusRoundCollecting, updateRound{ID: id})
}

func Collected(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundCollecting,
		player.PlayerRoundStatusRoundCollected, updateRound{ID: id})
}

func Submitting(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundCollected,
		player.PlayerRoundStatusRoundSubmitting, updateRound{ID: id})
}

func Submitted(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundSubmitting,
		player.PlayerRoundStatusRoundSubmitted, updateRound{ID: id})
}

func EndedJoined(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundSubmitted,
		player.PlayerRoundStatusRoundEnded, updateRound{ID: id})
}

func EndedExcluded(ctx context.Context, tx *sql.Tx, id int64) (rsql.NotifyFunc, error) {
	return fsm.UpdateTx(ctx, tx, player.PlayerRoundStatusRoundExcluded,
		player.PlayerRoundStatusRoundEnded, updateRound{ID: id})
}

func DeleteAll(ctx context.Context, dbc *sql.DB) error {
	tx, err := dbc.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "delete from player_rounds")
	if err != nil {
		return err
	}

	return tx.Commit()
}
