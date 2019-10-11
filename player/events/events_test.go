package events

import (
	"testing"

	"github.com/luno/reflex/rsql"

	"bitx/console/ticketing/internal/db"
)

func TestEventsTable(t *testing.T) {
	dbc := db.ConnectForTesting(t)
	defer dbc.Close()

	rsql.TestEventsTableInt(t, dbc, table)
}
