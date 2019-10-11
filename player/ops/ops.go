package ops

import "github.com/adamhicks/ctrlaltdefeat/player/config"

const TeamName = "CtrlAltDefeat"

func RunLoops(c config.Config, backends Backends) {
	// Add loops to run here!
	// e.g. go JoinRoundsForever()
	go StartMatchForever(backends, c)
	go ConsumeMatchEventsForever(backends, c)
	go StartRoundsForever(backends, c)
	go ConsumeRoundCollectEventsForever(c, backends)
	go JoinRoundsForever(c, backends)
	go CollectRoundsForever(c, backends)
	go SubmitRoundsForever(backends, c)
	go ConsumeRoundSubmitsForever(backends, c)
}
