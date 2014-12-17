package code

// This is a sample application which shows a simple use case of
// simulate. This usecase hits the Google homepage once a second -
// that's the default interval behavior.
// The simulate faults to writing to stdout/stderr.

import (
	"fmt"
	"net/url"

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
	resp, err := simhttp.Get("http://localhost:7676/jobs")
	if err != nil {
		return fmt.Errorf("got Error %+v ", err)
	}
	fmt.Printf("GET localhost:7676 - response time %f seconds. \n\n", resp.Duration.Seconds())
	return nil
}

func (m *MyActor) Post() error {
	resp, err := simhttp.Post("http://localhost:7676/jobs", url.Values{})
	if err != nil {
		return fmt.Errorf("got Error %+v ", err)
	}
	fmt.Printf("POST localhost:7676 - response time %f seconds. \n\n", resp.Duration.Seconds())
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