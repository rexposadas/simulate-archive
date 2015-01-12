package code

// This is a sample application which shows a simple use case of
// simulate. This usecase hits the Google homepage once a second -
// that's the default interval behavior.
// The simulate faults to writing to stdout/stderr.

import (
	"net/url"
	"time"

	"github.com/rexposadas/simulate"
	simhttp "github.com/rexposadas/simulate/http"
)

type MyActor struct{}

func (m *MyActor) Run() error {
	m.Get()
	m.Post()
	return nil
}

func (m *MyActor) Get() error {

	t := time.NewTicker(time.Second * 2)
	for {
		simhttp.Get("http://localhost:7676/jobs")
		<-t.C
	}

	return nil
}

func (m *MyActor) Post() error {

	t := time.NewTicker(time.Second)
	for {
		simhttp.Post("http://localhost:7676/jobs", url.Values{})
		<-t.C
	}
	return nil
}

func RunSimulator() {

	// The simulater is a service which makes API calls
	// no need to run simulate's REST endpoint for this example
	simulate.Run()

	// Create job and send to scheduler
	j := simulate.NewJob()
	a := &MyActor{}
	j.Actor = a
	simulate.Jobs <- j
}
