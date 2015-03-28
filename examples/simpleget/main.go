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
	"time"

	"github.com/franela/goreq"
	"github.com/rexposadas/simulate"
)

type MyActor struct{}

// GetGoogle make a GET request to http://google.com
func (m *MyActor) Run() error {
	t := time.NewTicker(time.Second)

	req := goreq.Request{
		Uri: "http://google.com",
	}
	for {
		s, err := simulate.MakeRequest(req)
		if err != nil {
			return fmt.Errorf("got Error %+v ", err)
		}
		s.Response.Body.Close()
		<-t.C
	}
	return nil
}

func main() {

	c := simulate.NewConfig()
	simulate.Run(c)

	// Create job and send to scheduler
	j := simulate.NewJob()
	d := &MyActor{}
	j.Actor = d
	simulate.Jobs <- j
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGQUIT)
	<-sigc
}
