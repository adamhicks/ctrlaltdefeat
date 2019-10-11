package roundsdb

import (
	"context"
	"database/sql"
	"github.com/adamhicks/ctrlaltdefeat/player"
)

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