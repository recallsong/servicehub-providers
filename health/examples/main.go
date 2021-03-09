package main

import (
	"os"

	"github.com/recallsong/servicehub"
	_ "github.com/recallsong/servicehub-providers/health"
	_ "github.com/recallsong/servicehub-providers/httpserver"
)

func main() {
	hub := servicehub.New()
	hub.Run("examples", "", os.Args...)
}
