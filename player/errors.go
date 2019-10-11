package player

import (
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
)

var (
	ErrPartNotCollected = errors.New("unable to find part for player", j.C("ERR_3f2c5232a172a25"))
	ErrNotInRound       = errors.New("player not in round", j.C("ERR_93759302f3252ac"))
)
