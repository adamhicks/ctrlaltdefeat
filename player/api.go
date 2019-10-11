package player

import (
	"context"

	"github.com/luno/reflex"
)

type Client interface {
	Stream(ctx context.Context, after string, opts ...reflex.StreamOption) (reflex.StreamClient, error)
	GetRoundParts(ctx context.Context, roundID int64) (RoundInfo, error)
}
