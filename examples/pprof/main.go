package main

import (
	"os"

	"github.com/recallsong/servicehub"
	_ "github.com/recallsong/servicehub-providers/providers/pprof"
)

func main() {
	hub := servicehub.New()
	hub.Run("examples", os.Args...)
}
