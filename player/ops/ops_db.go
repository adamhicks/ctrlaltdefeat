package ops

import (
	"context"
	"database/sql"

	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/adamhicks/ctrlaltdefeat/player/config"
	partsdb "github.com/adamhicks/ctrlaltdefeat/player/db/parts"
)

func GetPlayerParts(ctx context.Context, c config.Config, dbc *sql.DB, roundID int64) ([]player.RoundInfo, error) {
	parts, err := partsdb.GetRoundParts(ctx, dbc, int(roundID))
	if err != nil {
		return nil, err
	}
	p1 := c.GetPlayer(0)
	p2 := c.GetPlayer(1)
	p3 := c.GetPlayer(2)
	p4 := c.GetPlayer(3)

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
