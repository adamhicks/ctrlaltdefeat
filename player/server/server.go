package server

import (
	"context"

	"github.com/adamhicks/ctrlaltdefeat/player/events"
	"github.com/adamhicks/ctrlaltdefeat/player/ops"
	"github.com/adamhicks/ctrlaltdefeat/player/playerpb"
	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"
)

var _ playerpb.PlayerServer = (*Server)(nil)

// Server implements the engine grpc server.
type Server struct {
	b       ops.Backends
	rserver *reflex.Server
	stream  reflex.StreamFunc
}

// New returns a new server instance.
func New(b ops.Backends) *Server {
	return &Server{
		b:       b,
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

func (srv *Server) GetRoundParts(context.Context, *playerpb.GetRoundReq) (*playerpb.RoundInfo, error) {
	// todo: Implement a GetRoundParts
	return &playerpb.RoundInfo{}, nil
}
