package ops

import (
	"github.com/adamhicks/ctrlaltdefeat/player/config"
	"github.com/prometheus/common/log"
)

const TeamName = "CtrlAltDefeat"

func RunLoops(backends Backends, c config.Config) {
	// Add loops to run here!
	// e.g. go JoinRoundsForever()
	log.Info(nil, "Starting background loops")
	go StartMatchForever(backends, c)
	go ConsumeMatchEventsForever(backends, c)
	go StartRoundsForever(backends)
	go JoinRoundsForever(backends, c)
	go ConsumeRoundCollectEventsForever(backends)
	go CollectRoundsForever(backends, c)
	go SubmitRoundsForever(backends, c)
	go ConsumeRoundSubmitsForever(backends, c)
	go ConsumeMatchEndedForever(backends)
	go ConsumeRoundEndedForever(c, backends)
}
