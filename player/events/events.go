package events

import (
	"context"
	"database/sql"

	"github.com/adamhicks/ctrlaltdefeat/player"

	"github.com/luno/reflex"
	"github.com/luno/reflex/rsql"
)

const tableName = "player_rounds_events"

var table = rsql.NewEventsTableInt(tableName)

func Insert(ctx context.Context, tx *sql.Tx, foreignID int64,
	typ player.PlayerRoundStatus) (func(), error) {

	return table.Insert(ctx, tx, foreignID, typ)
}

func ToStream(dbc *sql.DB) reflex.StreamFunc {
	return table.ToStream(dbc)
}

func GetTable() rsql.EventsTableInt {
	return table
}
