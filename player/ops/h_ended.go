package ops

import (
	"context"

	"github.com/adamhicks/ctrlaltdefeat/player/db/cursors"
	partsdb "github.com/adamhicks/ctrlaltdefeat/player/db/parts"
	"github.com/adamhicks/ctrlaltdefeat/player/db/rounds"
	"github.com/corverroos/unsure"
	"github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/reflex"
)

const endCursor = "end_events"

//Listen for MatchEnded event, wipe db
func ConsumeMatchEndedForever(b Backends) {
	processMatchEvents := func(ctx context.Context, fate fate.Fate, event *reflex.Event) error {
		if !reflex.IsType(event.Type, engine.EventTypeMatchEnded) {
			return fate.Tempt()
		}
		if err := partsdb.DeleteAll(ctx, b.DB()); err != nil {
			return err
		}
		if err := rounds.DeleteAll(ctx, b.DB()); err != nil {
			return err
		}
		return fate.Tempt()
	}

	consumer := reflex.NewConsumer(endCursor, processMatchEvents)
	consumable := reflex.NewConsumable(b.EngineClient().Stream, cursors.ToStore(b.DB()))
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}
