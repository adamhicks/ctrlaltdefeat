package ops

import "github.com/adamhicks/ctrlaltdefeat/player/config"

const TeamName = "CtrlAltDefeat"

func RunLoops(c config.Config, backends Backends) {
	// Add loops to run here!
	// e.g. go JoinRoundsForever()
	go StartMatchForever(backends, c)
	go ConsumeMatchEventsForever(backends, c)
	go StartRoundsForever(backends)

	go JoinRoundsForever(backends, c)

	go ConsumeRoundCollectEventsForever(backends)

	go CollectRoundsForever(c, backends)
	go SubmitRoundsForever(backends, c)
	go ConsumeRoundSubmitsForever(backends, c)

	go ConsumeMatchEndedForever(c, backends)
}
