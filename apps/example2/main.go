package main

// This is a sample application which shows a simple use case of
// simulate. This usecase hits the Google homepage once a second -
// that's the default interval behavior.
// The server defaults to writing to stdout/stderr.

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rexposadas/simulate/apps/example2/code"
)

func main() {

	// start the API we will simulate against
	go code.StartAPI(7676)

	// run simulations against the our fake API
	go code.RunSimulator()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGQUIT)
	<-sigc
}
