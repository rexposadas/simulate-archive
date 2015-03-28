package fakeapi

// This is a sample application which shows a simple use case of
// simulate. This usecase hits the Google homepage once a second -
// that's the default interval behavior.
// The simulate faults to writing to stdout/stderr.

import (
	"time"

	"github.com/franela/goreq"
	"github.com/rexposadas/simulate"
)

type MyActor struct{}

func (m *MyActor) Run() error {
	m.Get()
	m.Post()
	return nil
}

func (m *MyActor) Get() error {

	t := time.NewTicker(time.Second * 2)
	req := goreq.Request{
		Uri: "http://localhost:7676/jobs",
	}
	for {
		s, err := simulate.MakeRequest(req)
		if err != nil {
			return err
		}
		defer s.Response.Body.Close()

		<-t.C
	}

	return nil
}

func (m *MyActor) Post() error {

	req := goreq.Request{
		Method: "POST",
		Uri:    "http://localhost:7676/jobs",
	}

	t := time.NewTicker(time.Second)
	for {
		s, err := simulate.MakeRequest(req)
		if err != nil {
			return err
		}
		defer s.Response.Body.Close()
		<-t.C
	}
	return nil
}

func RunSimulator() {

	// The simulater is a service which makes API calls
	// no need to run simulate's REST endpoint for this example
	c := simulate.NewConfig()
	simulate.Run(c)

	// Create job and send to scheduler
	j := simulate.NewJob()
	a := &MyActor{}
	j.Actor = a
	simulate.Jobs <- j
}
