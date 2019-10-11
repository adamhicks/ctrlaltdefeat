package rounds

import (
	"context"
	"database/sql"

	"github.com/adamhicks/ctrlaltdefeat/player"
)

// Call FSM functions here.

func Joining(ctx context.Context, dbc *sql.DB) (int64, error) {
	return fsm.Insert(ctx, dbc, joining{})
}

func Excluded(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, player.PlayerRoundStatusRoundJoining,
		player.PlayerRoundStatusRoundExcluded, excluded{ID: id})
}

func Joined(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, player.PlayerRoundStatusRoundJoining,
		player.PlayerRoundStatusRoundJoined, joined{ID: id})
}

func Collecting(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, player.PlayerRoundStatusRoundJoined,
		player.PlayerRoundStatusRoundCollecting, joined{ID: id})
}

func Collected(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, player.PlayerRoundStatusRoundCollecting,
		player.PlayerRoundStatusRoundCollected, joined{ID: id})
}

func Submitting(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, player.PlayerRoundStatusRoundCollected,
		player.PlayerRoundStatusRoundSubmitting, joined{ID: id})
}

func Submitted(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, player.PlayerRoundStatusRoundSubmitting,
		player.PlayerRoundStatusRoundSubmitted, joined{ID: id})
}

func EndedJoined(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, player.PlayerRoundStatusRoundSubmitted,
		player.PlayerRoundStatusRoundEnded, joined{ID: id})
}

func EndedExcluded(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, player.PlayerRoundStatusRoundExcluded,
		player.PlayerRoundStatusRoundEnded, joined{ID: id})
}
