package main

// This is a sample application which shows a simple use case of
// simulate. This usecase hits the Google homepage once a second -
// that's the default interval behavior.
// The simulate defaults to writing to stdout/stderr.

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rexposadas/simulate"
	simhttp "github.com/rexposadas/simulate/http"
)

type MyActor struct{}

// GetGoogle make a GET request to http://google.com
func (m *MyActor) Run() error {
	_, err := simhttp.Get("http://google.com")
	if err != nil {
		return fmt.Errorf("got Error %+v ", err)
	}
	return nil
}

func main() {

	simulate.Run()

	// Create job and send to scheduler
	j := simulate.NewJob()
	d := &MyActor{}
	j.Actor = d
	simulate.Jobs <- j
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGQUIT)
	<-sigc

	// stats check
	s := simulate.New()
	fmt.Printf("%+v\n", s.Count)
}
