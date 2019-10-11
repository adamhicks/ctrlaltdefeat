package rounds

import "github.com/adamhicks/ctrlaltdefeat/player"

//go:generate glean -table=player_rounds  -src=/Users/adam/unsure/

type glean struct {
	player.PlayerRound
}
