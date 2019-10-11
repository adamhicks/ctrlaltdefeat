package ops

import "github.com/adamhicks/ctrlaltdefeat/player/config"

const TeamName = "CtrlAltDefeat"

func RunLoops(c config.Config, backends Backends) {
	// Add loops to run here!
	// e.g. go JoinRoundsForever()
	go StartMatchForever(c, backends)
	go ConsumeMatchEventsForever(c, backends)
	go StartRoundsForever(c, backends)
	go ConsumeRoundCollectEventsForever(c, backends)
	go CollectRoundsForever(c, backends)

}
