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

func Excluded(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, player.PlayerRoundStatusRoundJoining,
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

func InsertRoundParts(ctx context.Context, dbc *sql.DB, matchId, roundId int, playerId string, rank, p1Part, p2Part, p3Part, p4Part int) (int64, error) {
	tx, err := dbc.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	r, err := tx.ExecContext(ctx, "insert into round_parts set match_id=?, round_id=?, player_id=?, rank=?, p1_part=?, p2_part=?, p3_part=?, p4_part=?", matchId, roundId, playerId, rank, p1Part, p2Part, p3Part, p4Part)
	if err != nil {
		return 0, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}

func GetRoundParts(ctx context.Context, dbc *sql.DB, roundId int, playerId int) (*player.RoundParts, error) {
	return scan(dbc.QueryRowContext(ctx, "select id, match_id, round_id, player_id, rank, p1_part, p2_part, p3_part, p4_part from round_parts where round_id=? and player_id=?", roundId, playerId))
}

func scan(r *sql.Row) (*player.RoundParts, error) {
	var roundParts player.RoundParts

	err := r.Scan(&roundParts.ID, &roundParts.MatchID, &roundParts.RoundID, &roundParts.Rank, &roundParts.P1Part, &roundParts.P2Part, &roundParts.P3Part, &roundParts.P4Part)
	if err != nil {
		return nil, err
	}

	return &roundParts, nil
}
