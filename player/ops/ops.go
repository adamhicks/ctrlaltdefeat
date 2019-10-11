package ops

import "github.com/adamhicks/ctrlaltdefeat/player/config"

const TeamName = "CtrlAltDefeat"

func RunLoops(c config.Config, backends Backends) {
	// Add loops to run here!
	go ConsumeRoundCollectEventsForever(c, backends)
	go JoinRoundsForever(c, backends)
	go CollectRoundsForever(c, backends)

}
