package main

import (
	"github.com/corverroos/unsure"
	"github.com/luno/jettison/log"
)

func main() {
	ctx := unsure.FatedContext()
	log.Info(ctx, "up and up")
}
