package ops

import (
	"context"
	"sort"

	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/adamhicks/ctrlaltdefeat/player/config"
	"github.com/adamhicks/ctrlaltdefeat/player/db/cursors"
	"github.com/adamhicks/ctrlaltdefeat/player/db/rounds"
	"github.com/corverroos/unsure"
	"github.com/luno/fate"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
	"github.com/pkg/errors"
)

func AmINext(c config.Config, parts []player.RoundInfo, playerJustBeen string) bool {
	me := c.GetMe()

	var mIdx, prevIdx = -1, -1

	for i, p := range parts {
		if p.Player == me.Name {
			mIdx = i
		}
		if p.Player == playerJustBeen {
			prevIdx = i
		}
	}
	if mIdx == 0 {
		return true
	}
	if mIdx == prevIdx+1 {
		return true
	}
	return false
}

//SubmitRoundsForever
//Get PRs with state RoundCollected, get the round parts, if it's first SubmitRound and transition to RoundSubmitted
//ConsumeRoundSubmitsForever
//Listen for PRRoundSubmitted, get the round parts, if it's next SubmitRound and transition to RoundSubmitted
func SubmitRoundsForever(b Backends, c config.Config) {
	ctx := unsure.FatedContext()
	for {
		collected, err := rounds.ListWithStatus(ctx, b.DB(), player.PlayerRoundStatusRoundCollected)
		if err != nil {
			log.Error(ctx, err)
			continue
		}

		for _, r := range collected {
			err := maybeGoToSubmitting(ctx, b, c, r, "")
			if err != nil {
				log.Error(ctx, err)
			}
		}
	}
}

func ConsumeRoundSubmitsForever(b Backends, c config.Config) {
	me := c.GetMe()
	for _, p := range c.GetAllPlayers() {
		if p == me {
			continue
		}
		ConsumePlayerRoundSubmitsForever(b, c, p)
	}
}

func ConsumePlayerRoundSubmitsForever(b Backends, c config.Config, p config.Player) {
	cli := b.GetPlayerClient(p.Name)
	var submitCursor = reflex.ConsumerName("submit_event_cursor_" + p.Name)

	f := func(ctx context.Context, fate fate.Fate, event *reflex.Event) error {
		if event.Type.ReflexType() != player.PlayerRoundStatusRoundSubmitted.ReflexType() {
			return nil
		}
		r, err := rounds.LookupRound(ctx, b.DB(), int(event.ForeignIDInt()))
		if err != nil {
			return err
		}
		err = maybeGoToSubmitting(ctx, b, c, r, p.Name)
		if err != nil {
			return err
		}
		return fate.Tempt()
	}

	consumer := reflex.NewConsumer(submitCursor, f)
	consumable := reflex.NewConsumable(cli.Stream, cursors.ToStore(b.DB()))
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}

func maybeGoToSubmitting(ctx context.Context, b Backends, c config.Config, r player.PlayerRound, prev string) error {
	parts, err := GetPlayerParts(ctx, c, b.DB(), r.RoundID)
	if err != nil {
		return errors.Wrap(err, "error getting parts")
	}
	if len(parts) != len(c.GetAllPlayers()) {
		return errors.New("Got incorrect number of parts")
	}

	sort.Slice(parts, func(i, j int) bool {
		r1, r2 := parts[i], parts[j]
		return r1.Rank < r2.Rank
	})

	if AmINext(c, parts, prev) {
		// Time to submit
		tx, err := b.DB().Begin()
		if err != nil {
			return err
		}
		notify, err := rounds.Submitting(ctx, tx, r.ID)
		if err != nil {
			return err
		}
		defer notify()
		return tx.Commit()
	}
	return nil
}
