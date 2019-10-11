package ops

import (
	"context"
	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/adamhicks/ctrlaltdefeat/player/config"
	"github.com/adamhicks/ctrlaltdefeat/player/db/cursors"
	"github.com/adamhicks/ctrlaltdefeat/player/db/rounds"
	"github.com/corverroos/unsure"
	"github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/reflex"
)

const endingCursor = "ending_events"

//Listen for EventTypeRoundSuccess and EventTypeRoundFailed, transition PR to RoundEnded
func ConsumeRoundEndedForever(c config.Config, b Backends) {
	processRoundEndedEvents := func(ctx context.Context, fate fate.Fate, event *reflex.Event) error {
		if !reflex.IsType(event.Type, engine.EventTypeRoundFailed) && !reflex.IsType(event.Type, engine.EventTypeRoundSuccess) {
			return fate.Tempt()
		}

		round, err := rounds.LookupRound(ctx, b.DB(), int(event.ForeignIDInt()))
		if err != nil {
			return err
		}

		tx, err := b.DB().Begin()
		if err != nil {
			return err
		}
		defer tx.Rollback()

		if int(round.Status) == player.PlayerRoundStatusRoundExcluded.Enum() {
			notify, err := rounds.EndedExcluded(ctx, tx, round.ID)
			if err != nil {
				return err
			}
			defer notify()
			return tx.Commit()
		} else if int(round.Status) == player.PlayerRoundStatusRoundSubmitted.Enum() {
			notify, err := rounds.EndedJoined(ctx, tx, round.ID)
			if err != nil {
				return err
			}
			defer notify()
			return tx.Commit()
		}

		return fate.Tempt()
	}

	consumer := reflex.NewConsumer(endingCursor, processRoundEndedEvents)
	consumable := reflex.NewConsumable(b.EngineClient().Stream, cursors.ToStore(b.DB()))
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}
