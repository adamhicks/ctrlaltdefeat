package state

import (
	"database/sql"
	"github.com/adamhicks/ctrlaltdefeat/player/db"
	"github.com/corverroos/unsure/engine"
	ec "github.com/corverroos/unsure/engine/client"
)

type State struct {
	engineClient engine.Client
	dbc               *sql.DB
}

func (s *State) EngineClient() engine.Client {
	return s.engineClient
}

func (s *State) DB() *sql.DB {
	return s.dbc
}

func New() (*State, error) {
	// TODO: Sergio: Get player id from somewhere
	p := "PLAYER OR INSTANCE ID HERE"
	dbc, err := db.Connect(p)
	if err != nil {
		return nil, err
	}

	engineClient, err := ec.New()
	if err != nil {
		return nil, err
	}

	return &State{
		engineClient: engineClient,
		dbc:            dbc.DB,
	}, nil
}
