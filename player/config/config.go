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

type Config struct {
	me  int
	all []Player
}

func New(idx int) Config {
	if idx < 0 || idx >= len(players) {
		panic(fmt.Sprintf("invalid player index %d", idx))
	}
	return Config{
		me:  idx,
		all: players,
	}
}

func (c Config) GetAllPlayers() []Player {
	return c.all
}

func (c Config) GetMe() Player {
	return c.all[c.me]
}
