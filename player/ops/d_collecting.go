package ops

import (
	"context"
	"database/sql"
	"time"

	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/log"

	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/adamhicks/ctrlaltdefeat/player/config"
	partsdb "github.com/adamhicks/ctrlaltdefeat/player/db/parts"
	"github.com/adamhicks/ctrlaltdefeat/player/db/rounds"
	"github.com/corverroos/unsure"
	"github.com/corverroos/unsure/engine"
)

//CollectRoundsForever
//Get PRs with state RoundCollecting, get the round parts, try and fetch missing data, or if complete transition to RoundCollected

func CollectRoundsForever(c config.Config, backends Backends) {
	dbc := backends.DB()
	ctx := unsure.FatedContext()
	ec := backends.EngineClient()
	me := c.GetMe().Name
	for {
		rs, err := rounds.ListWithStatus(ctx, dbc, player.PlayerRoundStatusRoundCollecting)
		if err != nil {
			log.Error(ctx, errors.Wrap(err, "Failed to list player_rounds in collecting state"), log.WithLevel(log.LevelInfo))
			continue
		}
		for _, r := range rs {
			res, err := ec.CollectRound(ctx, TeamName, me, r.RoundID)
			if err != nil {
				log.Error(ctx, errors.Wrap(err, "Failed to CollectRound"), log.WithLevel(log.LevelInfo))
				continue
			}
			if err = storeCollected(c, dbc, ctx, res, int(r.RoundID), me); err != nil {
				log.Error(ctx, errors.Wrap(err, "Failed to storeCollected"), log.WithLevel(log.LevelInfo))
				continue
			}
			tx, err := dbc.Begin()
			if err != nil {
				log.Error(ctx, errors.Wrap(err, "Failed to begin transaction"), log.WithLevel(log.LevelInfo))
				continue
			}
			if err = setCollected(ctx, tx, r.ID); err != nil {
				log.Error(ctx, errors.Wrap(err, "Failed setCollected"), log.WithLevel(log.LevelInfo))
				continue
			}

		}
		time.Sleep(time.Millisecond * 500)
	}
}

func storeCollected(c config.Config, dbc *sql.DB, ctx context.Context, res *engine.CollectRoundRes, roundID int, me string) error {
	rank := int64(res.Rank)
	p1Part, p2Part, p3Part, p4Part := 0, 0, 0, 0
	allPlayers := c.GetAllPlayers()
	for _, p := range res.Players {
		switch p.Name {
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
	_, err := partsdb.InsertRoundParts(ctx, dbc, roundID, me, rank, int64(p1Part), int64(p2Part), int64(p3Part), int64(p4Part))
	return err
}

func setCollected(ctx context.Context, tx *sql.Tx, id int64) error {
	defer tx.Rollback()
	notify, err := rounds.Collected(ctx, tx, id)
	if err != nil {
		return err
	}
	defer notify()
	return nil
}
