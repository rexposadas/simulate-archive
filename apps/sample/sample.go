package main

// This is a sample application which shows a simple use case of
// simulate. This usecase hits the Google homepage once a second -
// that's the default interval behavior.
// The server faults to writing to stdout/stderr.

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rexposadas/simulate/server"
)

func main() {

	server.Run(7000)

	// GET request every 1 second.
	server.Add("http://limitless-harbor-6554.herokuapp.com/")

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-sigc
}
