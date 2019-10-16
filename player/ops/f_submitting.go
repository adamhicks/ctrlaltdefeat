package ops

import (
	"context"

	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/adamhicks/ctrlaltdefeat/player/config"
	"github.com/adamhicks/ctrlaltdefeat/player/db/cursors"
	"github.com/adamhicks/ctrlaltdefeat/player/db/rounds"
	"github.com/corverroos/unsure"
	"github.com/luno/fate"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
)

// ConsumeSubmittingForever
// Listen for PlayerRoundStatusRoundSubmitting, SubmitRound to engine and transition to Submitted
func ConsumeSubmittingForever(b Backends, c config.Config) {
	me := c.GetMe() // Can only submit for myself
	cli := b.GetPlayerClient(me.Name)
	const submitCursor = "submitting_event_cursor"

	f := func(ctx context.Context, fate fate.Fate, event *reflex.Event) error {
		if event.Type.ReflexType() != player.PlayerRoundStatusRoundSubmitting.ReflexType() {
			return nil
		}
		r, err := rounds.LookupRound(ctx, b.DB(), int(event.ForeignIDInt()))
		if err != nil {
			return err
		}
		err = submitRound(ctx, b, c, r)
		if err != nil {
			log.Error(ctx, err)
		}
		return err
	}

	consumer := reflex.NewConsumer(submitCursor, f)
	consumable := reflex.NewConsumable(cli.Stream, cursors.ToStore(b.DB()))
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}

func submitRound(ctx context.Context, b Backends, c config.Config, r player.PlayerRound) error {
	me := c.GetMe()

	allParts, err := GetPlayerParts(ctx, c, b.DB(), r.RoundID)
	if err != nil {
		return err
	}
	total := CalcTotal(me, allParts)
	err = b.EngineClient().SubmitRound(ctx, c.GetTeam(), me.Name, r.RoundID, int(total))
	if err != nil {
		return err
	}
	tx, err := b.DB().Begin()
	if err != nil {
		return err
	}
	notify, err := rounds.Submitted(ctx, tx, r.ID)
	if err != nil {
		return err
	}
	defer notify()
	return tx.Commit()
}
