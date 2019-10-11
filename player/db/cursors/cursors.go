package cursors

import (
	"database/sql"

	"github.com/luno/reflex"
	"github.com/luno/reflex/rsql"
)

var cursors = rsql.NewCursorsTable("player_cursors")

// This returns the cursor store
func ToStore(dbc *sql.DB) reflex.CursorStore {
	return cursors.ToStore(dbc, rsql.WithCursorAsyncDisabled())
}
