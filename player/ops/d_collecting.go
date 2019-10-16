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
	me := c.GetMe()
	for _, p := range all {
		if !alreadyGot[p.Name] {
			var rInfo player.RoundInfo
			var err error
			if p == me {
				rInfo, err = collectFromEngine(ctx, b, c, r.RoundID)
			} else {
				rInfo, err = collectFromPlayer(ctx, b, c, p, r.RoundID)
			}
			if err != nil {
				log.Error(ctx, err)
				continue
			}
			err = storeRoundInfo(ctx, b, c, r.RoundID, rInfo)
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

func collectFromEngine(ctx context.Context, b Backends, c config.Config, roundID int64) (player.RoundInfo, error) {
	me := c.GetMe()
	res, err := b.EngineClient().CollectRound(ctx, c.GetTeam(), me.Name, roundID)
	if err != nil {
		return player.RoundInfo{}, err
	}

	var parts []player.PartInfo
	for _, p := range res.Players {
		parts = append(parts, player.PartInfo{Player: p.Name, Part: int64(p.Part)})
	}
	return player.RoundInfo{
		Player: me.Name,
		Rank:   int64(res.Rank),
		Parts:  parts,
	}, nil
}

func collectFromPlayer(ctx context.Context, b Backends, c config.Config, p config.Player, roundID int64) (player.RoundInfo, error) {
	cli := b.GetPlayerClient(p.Name)
	res, err := cli.GetRoundParts(ctx, roundID, c.GetMe().Name)
	if errors.Is(err, player.ErrNotInRound) {
		return player.RoundInfo{
			Player: p.Name,
			Rank:   -1,
			Parts:  nil,
		}, nil
	}
	return res, err
}

func storeRoundInfo(ctx context.Context, b Backends, c config.Config, roundID int64, round player.RoundInfo) error {
	rank := int64(round.Rank)
	var p1Part, p2Part, p3Part, p4Part int64 = 0, 0, 0, 0
	for _, p := range round.Parts {
		for i, cP := range c.GetAllPlayers() {
			if cP.Name == p.Player {
				switch i {
				case 0:
					p1Part = p.Part
					break
				case 1:
					p2Part = p.Part
					break
				case 2:
					p3Part = p.Part
					break
				case 3:
					p4Part = p.Part
					break
				}
			}
		}
	}
	_, err := partsdb.InsertRoundParts(
		ctx, b.DB(), int(roundID), round.Player,
		rank, p1Part, p2Part, p3Part, p4Part,
	)
	return err
}
