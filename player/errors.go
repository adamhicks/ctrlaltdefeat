package player

import (
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
)

var (
	ErrPlayerNotFound = errors.New("unable to find part for player", j.C("ERR_3f2c5232a172a25"))
)
