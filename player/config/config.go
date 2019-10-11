package config

import "fmt"

type Player struct {
	Name     string
	GRPCPort int
}

var players = []Player{
	{Name: "wes", GRPCPort: 17513},
	{Name: "tom", GRPCPort: 17514},
	{Name: "sergio", GRPCPort: 17515},
	{Name: "adam", GRPCPort: 17516},
}

func GetPlayer(idx int) Player {
	if idx < 0 || idx >= len(players) {
		panic(fmt.Sprintf("invalid player index %d", idx))
	}
	return players[idx]
}

func GetAllPlayers() []Player {
	return players
}
