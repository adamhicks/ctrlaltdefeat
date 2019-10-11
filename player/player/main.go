package main

import (
	"flag"

	"github.com/adamhicks/ctrlaltdefeat/player/config"
	"github.com/adamhicks/ctrlaltdefeat/player/state"
	"github.com/corverroos/unsure"
	"github.com/luno/jettison/errors"
)

var playerIdx = flag.Int("player_idx", 0, "the index into player array for this instance")

func main() {
	unsure.Bootstrap()

	p := config.GetPlayer(*playerIdx)

	_, err := state.New(p, config.GetAllPlayers())
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new state error"))
	}

	unsure.WaitForShutdown()
}
