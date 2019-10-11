package state

import (
	"database/sql"
	"fmt"

	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/adamhicks/ctrlaltdefeat/player/client"
	"github.com/adamhicks/ctrlaltdefeat/player/config"
	"github.com/adamhicks/ctrlaltdefeat/player/db"
	"github.com/corverroos/unsure/engine"
	ec "github.com/corverroos/unsure/engine/client"
)

type State struct {
	engineClient  engine.Client
	dbc           *sql.DB
	playerClients map[string]player.Client
}

func (s *State) EngineClient() engine.Client {
	return s.engineClient
}

func (s *State) DB() *sql.DB {
	return s.dbc
}

func (s *State) GetPlayerClient(playerName string) player.Client {
	c, ok := s.playerClients[playerName]
	if !ok {
		panic("Unknown player '" + playerName + "'")
	}
	return c
}

func New(c config.Config) (*State, error) {
	me := c.GetMe()
	all := c.GetAllPlayers()

	dbc, err := db.Connect(me.Name)
	if err != nil {
		return nil, err
	}

	engineClient, err := ec.New()
	if err != nil {
		return nil, err
	}

	playerClients := make(map[string]player.Client, len(all))
	for _, p := range all {
		c, err := client.New(client.WithAddress(fmt.Sprintf("localhost:%d", p.GRPCPort)))
		if err != nil {
			return nil, err
		}
		playerClients[p.Name] = c
	}

	return &State{
		engineClient:  engineClient,
		dbc:           dbc,
		playerClients: playerClients,
	}, nil
}
