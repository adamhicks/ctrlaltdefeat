package ops

import (
	"context"
	"database/sql"

	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/adamhicks/ctrlaltdefeat/player/config"
	partsdb "github.com/adamhicks/ctrlaltdefeat/player/db/parts"
)

func GetPlayerPart(
	ctx context.Context, c config.Config, dbc *sql.DB, roundID int64, playerName string,
) (player.RoundInfo, error) {
	parts, err := GetPlayerParts(ctx, c, dbc, roundID)
	if err != nil {
		return player.RoundInfo{}, err
	}
	me := c.GetMe()
	for _, p := range parts {
		if p.Player == me.Name {
			return p, nil
		}
	}
	return player.RoundInfo{}, player.ErrPartNotCollected
}

func GetPlayerParts(ctx context.Context, c config.Config, dbc *sql.DB, roundID int64) ([]player.RoundInfo, error) {
	parts, err := partsdb.GetRoundParts(ctx, dbc, int(roundID))
	if err != nil {
		return nil, err
	}
	var p1, p2, p3, p4 config.Player
	if c.GetPlayerCount() > 0 {
		p1 = c.GetPlayer(0)
	}
	if c.GetPlayerCount() > 1 {
		p2 = c.GetPlayer(1)
	}
	if c.GetPlayerCount() > 2 {
		p3 = c.GetPlayer(2)
	}
	if c.GetPlayerCount() > 3 {
		p4 = c.GetPlayer(3)
	}

	var rounds []player.RoundInfo
	for _, p := range parts {
		r := player.RoundInfo{
			Player: p.PlayerID,
			Rank:   p.Rank,
			Parts: []player.PartInfo{
				{Player: p1.Name, Part: p.P1Part},
				{Player: p2.Name, Part: p.P2Part},
				{Player: p3.Name, Part: p.P3Part},
				{Player: p4.Name, Part: p.P4Part},
			},
		}
		rounds = append(rounds, r)
	}
	return rounds, nil
}

func CalcTotal(p config.Player, roundInfos []player.RoundInfo) int64 {
	var sum int64
	for _, info := range roundInfos {
		for _, part := range info.Parts {
			if part.Player == p.Name {
				sum += part.Part
				break
			}
		}
	}
	return sum
}
