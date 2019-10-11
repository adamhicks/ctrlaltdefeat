package ops

import (
	"context"
	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/adamhicks/ctrlaltdefeat/player/config"
	"github.com/adamhicks/ctrlaltdefeat/player/db/cursors"
	"github.com/adamhicks/ctrlaltdefeat/player/db/rounds"
	"github.com/corverroos/unsure"
	"github.com/corverroos/unsure/engine"
	_ "github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/reflex"
)

//JoinRoundsForever
//Get PRs with state RoundJoining and call JoinRound if true go to RoundJoined, if false go to RoundExcluded

const roundJoiningConsumer = "round_joining"

func JoinRoundsForever(p config.Config, b Backends) {
	f := func(ctx context.Context, fate fate.Fate, e *reflex.Event) error {
		if !reflex.IsType(e.Type, engine.EventTypeRoundJoin) {
			return fate.Tempt()
		}

		pr, err := rounds.LookupRoundAndStatus(ctx, b.DB(), e.ForeignIDInt(), player.PlayerRoundStatusRoundJoining)
		if err != nil {
			return fate.Tempt()
		}

		joined, err := b.EngineClient().JoinRound(ctx, TeamName, p.GetMe().Name, pr.RoundID)
		if err != nil {
			return err
		}

		tx, err := b.DB().Begin()
		if err != nil {
			return err
		}

		if joined {
			_, err := rounds.Joined(ctx, tx, pr.RoundID)
			if err != nil {
				return err
			}

			return fate.Tempt()
		}

		_, err = rounds.Excluded(ctx, tx, pr.RoundID)
		if err != nil {
			return err
		}

		return fate.Tempt()
	}

	cli := b.GetPlayerClient(p.GetMe().Name)
	consumer := reflex.NewConsumer(roundJoiningConsumer, f)
	consumable := reflex.NewConsumable(cli.Stream, cursors.ToStore(b.DB()))
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}