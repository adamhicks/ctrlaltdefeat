package ops

import (
	"context"
	"time"

	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/adamhicks/ctrlaltdefeat/player/config"
	partsdb "github.com/adamhicks/ctrlaltdefeat/player/db/parts"
	"github.com/adamhicks/ctrlaltdefeat/player/db/rounds"
	"github.com/corverroos/unsure"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/log"
)

//CollectRoundsForever
//Get PRs with state RoundCollecting, get the round parts, try and fetch missing data, or if complete transition to RoundCollected

func CollectRoundsForever(b Backends, c config.Config) {
	for {
		ctx := unsure.FatedContext()
		rs, err := rounds.ListWithStatus(ctx, b.DB(), player.PlayerRoundStatusRoundCollecting)
		if err != nil {
			log.Error(ctx, err)
			continue
		}
		for _, r := range rs {
			err := collectRoundParts(ctx, b, c, r)
			if err != nil {
				log.Error(ctx, err)
				continue
			}
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func collectRoundParts(ctx context.Context, b Backends, c config.Config, r player.PlayerRound) error {
	parts, err := GetPlayerParts(ctx, c, b.DB(), r.RoundID)
	if err != nil {
		return err
	}
	alreadyGot := make(map[string]bool)
	for _, p := range parts {
		alreadyGot[p.Player] = true
	}
	all := c.GetAllPlayers()
	for _, p := range all {
		if !alreadyGot[p.Name] {
			err := fetchAndStorePart(ctx, b, c, p, r.RoundID)
			if err != nil {
				log.Error(ctx, err)
				continue
			}
			alreadyGot[p.Name] = true
		}
	}
	if len(alreadyGot) == len(all) {
		tx, err := b.DB().Begin()
		if err != nil {
			return err
		}
		notify, err := rounds.Collected(ctx, tx, r.ID)
		if err != nil {
			return err
		}
		defer notify()
		return tx.Commit()
	}
	return nil
}

func fetchAndStorePart(ctx context.Context, b Backends, c config.Config, p config.Player, roundID int64) error {
	cli := b.GetPlayerClient(p.Name)
	res, err := cli.GetRoundParts(ctx, roundID, c.GetMe().Name)
	if errors.Is(err, player.ErrNotInRound) {
		_, err := partsdb.InsertRoundParts(
			ctx, b.DB(), int(roundID), p.Name,
			-1, 0, 0, 0, 0,
		)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	rank := int64(res.Rank)
	var p1Part, p2Part, p3Part, p4Part int64 = 0, 0, 0, 0
	allPlayers := c.GetAllPlayers()
	for _, p := range res.Parts {
		switch p.Player {
		case allPlayers[0].Name:
			p1Part = p.Part
		case allPlayers[1].Name:
			p2Part = p.Part
		case allPlayers[2].Name:
			p3Part = p.Part
		case allPlayers[3].Name:
			p4Part = p.Part
		}
	}
	_, err = partsdb.InsertRoundParts(
		ctx, b.DB(), int(roundID), p.Name,
		rank, p1Part, p2Part, p3Part, p4Part,
	)
	return err
}
