package cursors

import (
	"testing"

	"github.com/adamhicks/ctrlaltdefeat/player/db"
	"github.com/corverroos/unsure"
	"github.com/luno/reflex/rsql"
)

func TestCursorsTable(t *testing.T) {
	defer unsure.CheatFateForTesting(t)()
	dbc := db.ConnectForTesting(t)
	defer dbc.Close()

	rsql.TestCursorsTable(t, dbc, cursors)
}
