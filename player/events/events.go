package events

import (
	"context"
	"database/sql"

	"github.com/adamhicks/ctrlaltdefeat/player"

	"github.com/luno/reflex"
	"github.com/luno/reflex/rsql"
)

const tableName = "player_rounds_events"

var table = rsql.NewEventsTableInt(tableName,
	rsql.WithEventsCacheEnabled(),
	rsql.WithEventsNotifier(beepwrap.New(tableName)))

// Insert creates a attachment_upload_events record in the db.
func Insert(ctx context.Context, tx *sql.Tx, foreignID int64,
	typ player.EventType) (func(), error) {

	return table.Insert(ctx, tx, foreignID, typ)
}

// ToStream returns a reflex stream for attachment update events.
func ToStream(dbc *sql.DB) reflex.StreamFunc {
	return table.ToStream(dbc)
}

//GetTable returns a reflex events table.
func GetTable() rsql.EventsTableInt {
	return table
}
