package main

// This is a sample application which shows a simple use case of
// simulate. This usecase hits the Google homepage once a second -
// that's the default interval behavior.
// The server faults to writing to stdout/stderr.

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	simhttp "github.com/rexposadas/simulate/http"
	"github.com/rexposadas/simulate/server"
)

// GetGoogle make a GET request to http://google.com
func GetGoogle() error {
	resp, err := simhttp.Get("http://google.com")
	if err != nil {
		return fmt.Errorf("got Error %+v ", err)
	}
	fmt.Printf("GetGoogle - response time %f seconds. \n\n", resp.Duration.Seconds())
	return nil
}

func main() {

	server.Run(0)

	// Create job and send to scheduler
	g := server.NewJob()
	g.Run = GetGoogle

	server.Jobs <- g

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGQUIT)
	<-sigc
}
