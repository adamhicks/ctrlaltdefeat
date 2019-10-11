package ops

import "github.com/adamhicks/ctrlaltdefeat/player/config"

const TeamName = "CtrlAltDefeat"

func RunLoops(c config.Config, backends Backends) {
	// Add loops to run here!
	// e.g. go JoinRoundsForever()
	go ConsumeRoundCollectEventsForever(c, backends)
}
