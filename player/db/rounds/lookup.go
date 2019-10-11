package rounds

import (
	"context"
	"database/sql"

	"github.com/adamhicks/ctrlaltdefeat/player"
)

func ListWithStatus(ctx context.Context, dbc *sql.DB, s player.PlayerRoundStatus) ([]player.PlayerRound, error) {
	rows, err := dbc.QueryContext(ctx, selectPrefix+"status=?", s.Enum())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rounds []player.PlayerRound
	for rows.Next() {
		r, err := scan(rows)
		if err != nil {
			return nil, err
		}
		rounds = append(rounds, *r)
	}
	return rounds, nil
}

func LookupRound(ctx context.Context, dbc dbc, roundID int) (player.PlayerRound, error) {
	r, err := lookupWhere(ctx, dbc, "round_id=? order by id desc limit 1", roundID)
	if err != nil {
		return player.PlayerRound{}, err
	}
	return *r, nil
}

func LookupRoundAndStatus(ctx context.Context, dbc dbc, roundID int64, s player.PlayerRoundStatus) (player.PlayerRound, error) {
	r, err := lookupWhere(ctx, dbc, "round_id=? and status=?", roundID, s)
	if err != nil {
		return player.PlayerRound{}, err
	}
	return *r, nil
}
