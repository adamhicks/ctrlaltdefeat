package server

import (
	"context"

	"github.com/adamhicks/ctrlaltdefeat/player/config"
	"github.com/adamhicks/ctrlaltdefeat/player/events"
	"github.com/adamhicks/ctrlaltdefeat/player/ops"
	"github.com/adamhicks/ctrlaltdefeat/player/playerpb"
	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"
)

var _ playerpb.PlayerServer = (*Server)(nil)

// Server implements the player grpc server.
type Server struct {
	b       ops.Backends
	c       config.Config
	rserver *reflex.Server
	stream  reflex.StreamFunc
}

// New returns a new server instance.
func New(b ops.Backends, c config.Config) *Server {
	return &Server{
		b:       b,
		c:       c,
		rserver: reflex.NewServer(),
		stream:  events.ToStream(b.DB()),
	}
}

func (srv *Server) Stop() {
	srv.rserver.Stop()
}

func (srv *Server) Stream(req *reflexpb.StreamRequest, ss playerpb.Player_StreamServer) error {
	return srv.rserver.Stream(srv.stream, req, ss)
}

func (srv *Server) GetRoundParts(ctx context.Context, req *playerpb.GetRoundReq) (*playerpb.RoundInfo, error) {
	round, err := ops.GetPlayerPart(ctx, srv.c, srv.b.DB(), req.RoundId, req.Player)
	if err != nil {
		return nil, err
	}
	return playerpb.PlayerRoundInfoToProto(round), nil
}
