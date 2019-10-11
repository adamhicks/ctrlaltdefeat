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
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
	"time"
)

const startCursor = "start_events"

//Check for number of rounds not in RoundEnded state, if == 0, try to start a match
func StartMatchForever(config config.Config, b Backends) {
	for {
		ongoingRounds, err := rounds.ListWithStatusNot(unsure.FatedContext(), *b.DB(), player.PlayerRoundStatusRoundEnded)
		if err != nil {
			return
		}
		if len(ongoingRounds) == 0 {
			ctx := unsure.ContextWithFate(context.Background(), unsure.DefaultFateP())

			err := b.EngineClient().StartMatch(ctx, TeamName, len(config.GetAllPlayers()))

			if errors.Is(err, engine.ErrActiveMatch) {
				// Match active, just ignore
				return
			} else if err != nil {
				log.Error(ctx, errors.Wrap(err, "start match error"))
			} else {
				log.Info(ctx, "match started")
				return
			}
		}

		time.Sleep(time.Second)
	}
}

//Listen for MatchEnded event, try to start a match
func ConsumeMatchEventsForever(config config.Config, b Backends) {
	processMatchEvents := func (ctx context.Context, fate fate.Fate, event *reflex.Event) error {
		if !reflex.IsType(event.Type, engine.EventTypeMatchEnded) {
		return fate.Tempt()
	}

		return b.EngineClient().StartMatch(ctx, TeamName, len(config.GetAllPlayers()))
	}

	consumer := reflex.NewConsumer(startCursor, processMatchEvents)
	consumable := reflex.NewConsumable(b.EngineClient().Stream, cursors.ToStore(b.DB()))
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}


//Listen for EventTypeRoundJoin event and create a PlayerRound(PR) object
func StartRoundsForever(config config.Config, b Backends) {
	processRoundJoinEvents := func(ctx context.Context, fate fate.Fate, event *reflex.Event) error {
		if !reflex.IsType(event.Type, engine.EventTypeRoundJoin) {
			return fate.Tempt()
		}

		_, err := rounds.Joining(ctx, b.DB(), event.ForeignIDInt())
		if err != nil {
			return err
		}

		return fate.Tempt()
	}

	consumer := reflex.NewConsumer(startCursor, processRoundJoinEvents)
	consumable := reflex.NewConsumable(b.EngineClient().Stream, cursors.ToStore(b.DB()))
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}
