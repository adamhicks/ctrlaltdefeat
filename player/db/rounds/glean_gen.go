// Code generated by glean from lookup.go:5. DO NOT EDIT.
package rounds

import (
	"context"
	"database/sql"

	"github.com/adamhicks/ctrlaltdefeat/player"
)

const cols = " `id`, `round_id`, `status`, `created_at`, `updated_at` "
const selectPrefix = "select " + cols + " from player_rounds where "

func Lookup(ctx context.Context, dbc dbc, id int64) (*player.PlayerRound, error) {
	return lookupWhere(ctx, dbc, "id=?", id)
}

// lookupWhere queries the player_rounds table with the provided where clause, then scans
// and returns a single row.
func lookupWhere(ctx context.Context, dbc dbc, where string, args ...interface{}) (*player.PlayerRound, error) {
	return scan(dbc.QueryRowContext(ctx, selectPrefix+where, args...))
}

// listWhere queries the player_rounds table with the provided where clause, then scans
// and returns all the rows.
func listWhere(ctx context.Context, dbc dbc, where string, args ...interface{}) ([]player.PlayerRound, error) {

	rows, err := dbc.QueryContext(ctx, selectPrefix+where, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []player.PlayerRound
	for rows.Next() {
		r, err := scan(rows)
		if err != nil {
			return nil, err
		}
		res = append(res, *r)
	}

	return res, rows.Err()
}

func scan(row row) (*player.PlayerRound, error) {
	var g glean

	err := row.Scan(&g.ID, &g.RoundID, &g.Status, &g.CreatedAt, &g.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &player.PlayerRound{
		ID:        g.ID,
		RoundID:   g.RoundID,
		Status:    g.Status,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}, nil
}

// dbc is a common interface for *sql.DB and *sql.Tx.
type dbc interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// row is a common interface for *sql.Rows and *sql.Row.
type row interface {
	Scan(dest ...interface{}) error
}