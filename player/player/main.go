package main

import (
	"flag"
	"fmt"

	"github.com/adamhicks/ctrlaltdefeat/player/config"
	"github.com/adamhicks/ctrlaltdefeat/player/ops"
	"github.com/adamhicks/ctrlaltdefeat/player/playerpb"
	"github.com/adamhicks/ctrlaltdefeat/player/server"
	"github.com/adamhicks/ctrlaltdefeat/player/state"
	"github.com/corverroos/unsure"
	"github.com/luno/jettison/errors"
)

var playerIdx = flag.Int("player_idx", 0, "the index into player array for this instance")

func main() {
	unsure.Bootstrap()

	c := config.New(*playerIdx)

	s, err := state.New(c)
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new state error"))
	}

	serveGRPCForever(c.GetMe(), s)
	ops.RunLoops(c, s)

	unsure.WaitForShutdown()
}

func serveGRPCForever(p config.Player, s *state.State) {
	addr := fmt.Sprintf("localhost:%d", p.GRPCPort)
	grpcServer, err := unsure.NewServer(addr)
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new server"))
	}

	pServer := server.New(s)
	playerpb.RegisterPlayerServer(grpcServer.GRPCServer(), pServer)

	unsure.RegisterNoErr(func() {
		pServer.Stop()
		grpcServer.Stop()
	})

	unsure.Fatal(grpcServer.ServeForever())
}
