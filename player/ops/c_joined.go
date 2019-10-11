package ops

import (
	"context"

	"github.com/adamhicks/ctrlaltdefeat/player/config"
	"github.com/adamhicks/ctrlaltdefeat/player/db/cursors"
	"github.com/corverroos/unsure"
	"github.com/luno/fate"
	"github.com/luno/reflex"
	"github.com/prometheus/common/log"
)

const joinedCursor = "joined_events"

//ConsumeRoundCollectEventsForever
//Listen for EventTypeRoundCollect, get the PR, transition to RoundCollecting
func ConsumeRoundCollectEventsForever(c config.Config, b Backends) {
	for _, p := range c.GetAllPlayers() {
		if p == c.GetMe() {
			continue
		}
		go consumePlayerCollects(p.Name, b)
	}
}

func consumePlayerCollects(p string, b Backends) {
	cli := b.GetPlayerClient(p)
	consumer := reflex.NewConsumer(joinedCursor, processRoundEvents)
	consumable := reflex.NewConsumable(cli.Stream, cursors.ToStore(b.DB()))
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}

func processRoundEvents(ctx context.Context, fate fate.Fate, event *reflex.Event) error {
	log.Info("Got an event")
	return nil
}
