// Code generated by glean from glean.go:7. DO NOT EDIT.
package rounds

import (
	"context"
	"database/sql"

	partsdb "github.com/adamhicks/ctrlaltdefeat/player/db/parts"
)

const cols = " `id`, `round_id`, `player_id`, `rank`, `p1_part`, `p2_part`, `p3_part`, `p4_part` "
const selectPrefix = "select " + cols + " from player_rounds where "

func Lookup(ctx context.Context, dbc dbc, id int64) (*partsdb.RoundParts, error) {
	return lookupWhere(ctx, dbc, "id=?", id)
}

// lookupWhere queries the player_rounds table with the provided where clause, then scans
// and returns a single row.
func lookupWhere(ctx context.Context, dbc dbc, where string, args ...interface{}) (*partsdb.RoundParts, error) {
	return scan(dbc.QueryRowContext(ctx, selectPrefix+where, args...))
}

// listWhere queries the player_rounds table with the provided where clause, then scans
// and returns all the rows.
func listWhere(ctx context.Context, dbc dbc, where string, args ...interface{}) ([]partsdb.RoundParts, error) {

	rows, err := dbc.QueryContext(ctx, selectPrefix+where, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []partsdb.RoundParts
	for rows.Next() {
		r, err := scan(rows)
		if err != nil {
			return nil, err
		}
		res = append(res, *r)
	}

	return res, rows.Err()
}

func scan(row row) (*partsdb.RoundParts, error) {
	var g glean

	err := row.Scan(&g.ID, &g.RoundID, &g.PlayerID, &g.Rank, &g.P1Part, &g.P2Part, &g.P3Part, &g.P4Part)
	if err != nil {
		return nil, err
	}

	return &partsdb.RoundParts{
		ID:       g.ID,
		RoundID:  g.RoundID,
		PlayerID: g.PlayerID,
		Rank:     g.Rank,
		P1Part:   g.P1Part,
		P2Part:   g.P2Part,
		P3Part:   g.P3Part,
		P4Part:   g.P4Part,
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