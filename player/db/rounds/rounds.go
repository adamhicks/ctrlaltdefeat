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

func GetRoundParts(ctx context.Context, dbc *sql.DB, roundId int, playerId int) (*[]partsdb.RoundParts, error) {
	rows, err := dbc.QueryContext(ctx, "select id, round_id, player_id, rank, p1_part, p2_part, p3_part, p4_part from round_parts where round_id=? and player_id=?", roundId, playerId)
	if err != nil {
		return nil, err
	}

	roundParts, err := scan(rows)
	if err != nil {
		return nil, err
	}

	return roundParts, nil
}

func scan(r *sql.Rows) (*[]partsdb.RoundParts, error) {
	var roundPartsArr []partsdb.RoundParts

	for r.Next() {
		var roundParts partsdb.RoundParts
		err := r.Scan(&roundParts.ID, &roundParts.RoundID, &roundParts.Rank, &roundParts.P1Part, &roundParts.P2Part, &roundParts.P3Part, &roundParts.P4Part)
		if err != nil {
			return nil, err
		}
		roundPartsArr = append(roundPartsArr, roundParts)
	}

	return &roundPartsArr, nil
}
