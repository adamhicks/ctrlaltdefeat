package roundsdb

import (
	"context"
	"database/sql"
	partsdb "github.com/adamhicks/ctrlaltdefeat/player/db/parts"
)

func InsertRoundParts(ctx context.Context, dbc *sql.DB, roundId int, playerId string, rank, p1Part, p2Part, p3Part, p4Part int64) (int64, error) {
	tx, err := dbc.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	r, err := tx.ExecContext(ctx, "insert into round_parts set round_id=?, player_id=?, rank=?, p1_part=?, p2_part=?, p3_part=?, p4_part=?", roundId, playerId, rank, p1Part, p2Part, p3Part, p4Part)
	if err != nil {
		return 0, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}

func GetRoundParts(ctx context.Context, dbc *sql.DB, roundId int, playerId int) ([]partsdb.RoundParts, error) {
	return listWhere(ctx, dbc, "round_id=?", roundId)
}
