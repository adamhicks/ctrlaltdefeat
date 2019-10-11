package roundsdb

import (
	partsdb "github.com/adamhicks/ctrlaltdefeat/player/db/parts"
)

//go:generate glean -table=player_rounds -src=/Users/wesley/work/

type glean struct {
	partsdb.RoundParts
}

