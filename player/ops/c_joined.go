package ops

import (
	"context"

	"github.com/adamhicks/ctrlaltdefeat/player/db/cursors"
	"github.com/adamhicks/ctrlaltdefeat/player/db/rounds"
	"github.com/corverroos/unsure"
	"github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
)

const joinedCursor = "joined_events"

//ConsumeRoundCollectEventsForever
//Listen for EventTypeRoundCollect, get the PR, transition to RoundCollecting
func ConsumeRoundCollectEventsForever(b Backends) {
	cli := b.EngineClient()

	f := func(ctx context.Context, fate fate.Fate, event *reflex.Event) error {
		if event.Type.ReflexType() != engine.EventTypeRoundCollect.ReflexType() {
			return nil
		}
		r, err := rounds.LookupRound(ctx, b.DB(), int(event.ForeignIDInt()))
		if err != nil {
			return err
		}
		log.Info(nil, "got status", j.KV("status", r.Status))
		tx, err := b.DB().Begin()
		if err != nil {
			return err
		}
		defer tx.Rollback()
		notify, err := rounds.Collecting(ctx, tx, r.ID)
		if err != nil {
			return err
		}
		defer notify()
		return tx.Commit()
	}

	consumer := reflex.NewConsumer(joinedCursor, f)
	consumable := reflex.NewConsumable(cli.Stream, cursors.ToStore(b.DB()))
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}
