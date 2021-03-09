package main

import (
	"os"

	"github.com/recallsong/servicehub"
	_ "github.com/recallsong/servicehub-providers/httpserver"
	_ "github.com/recallsong/servicehub-providers/pprof"
)

func main() {
	hub := servicehub.New()
	hub.Run("examples", "", os.Args...)
}
