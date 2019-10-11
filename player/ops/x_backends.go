package ops

import (
	"database/sql"

	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/corverroos/unsure/engine"
)

type Backends interface {
	DB() *sql.DB
	EngineClient() engine.Client
	GetPlayerClient(playerName string) player.Client
}
