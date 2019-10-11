package client

import (
	"context"
	"flag"

	"github.com/adamhicks/ctrlaltdefeat/player"
	"github.com/adamhicks/ctrlaltdefeat/player/playerpb"
	"github.com/corverroos/unsure"
	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"
	"google.golang.org/grpc"
)

var addr = flag.String("player_address", "", "host:port of player gRPC service")

var _ player.Client = (*client)(nil)

type client struct {
	address   string
	rpcConn   *grpc.ClientConn
	rpcClient playerpb.PlayerClient
}

type Option func(*client)

func WithAddress(address string) Option {
	return func(c *client) {
		c.address = address
	}
}

func New(opts ...Option) (*client, error) {
	c := client{
		address: *addr,
	}
	for _, o := range opts {
		o(&c)
	}

	var err error
	c.rpcConn, err = unsure.NewClient(c.address)
	if err != nil {
		return nil, err
	}

	c.rpcClient = playerpb.NewPlayerClient(c.rpcConn)

	return &c, nil
}

func (c *client) Stream(ctx context.Context, after string, opts ...reflex.StreamOption) (reflex.StreamClient, error) {
	sFn := reflex.WrapStreamPB(func(ctx context.Context,
		req *reflexpb.StreamRequest) (reflex.StreamClientPB, error) {
		return c.rpcClient.Stream(ctx, req)
	})
	return sFn(ctx, after, opts...)
}

func (c *client) GetRoundParts(ctx context.Context, roundID int64) (player.RoundInfo, error) {
	res, err := c.rpcClient.GetRoundParts(ctx, &playerpb.GetRoundReq{
		RoundId: roundID,
	})
	if err != nil {
		return player.RoundInfo{}, err
	}
	return playerpb.PlayerRoundInfoFromProto(res), nil
}
