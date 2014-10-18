package code

// This is a sample application which shows a simple use case of
// simulate. This usecase hits the Google homepage once a second -
// that's the default interval behavior.
// The server faults to writing to stdout/stderr.

import (
	"fmt"
	"net/url"

	simhttp "github.com/rexposadas/simulate/http"
	"github.com/rexposadas/simulate/server"
)

// GetGoogle make a GET request to http://google.com
func Get() error {
	resp, err := simhttp.Get("http://localhost:7676/jobs")
	if err != nil {
		return fmt.Errorf("got Error %+v ", err)
	}
	fmt.Printf("GET localhost:7676 - response time %f seconds. \n\n", resp.Duration.Seconds())
	return nil
}

func Post() error {
	resp, err := simhttp.Post("http://localhost:7676/jobs", url.Values{})
	if err != nil {
		return fmt.Errorf("got Error %+v ", err)
	}
	fmt.Printf("POST localhost:7676 - response time %f seconds. \n\n", resp.Duration.Seconds())
	return nil
}

func RunSimulator(port int) {

	// The simulater is a service which makes API calls
	// no need to run simulate's REST endpoint for this example
	server.Run(0)

	// Create job and send to scheduler
	g := server.NewJob()
	g.Run = Get
	server.Jobs <- g
	fmt.Println("added GET job")

	p := server.NewJob()
	p.Run = Post
	server.Jobs <- p
	fmt.Println("added POST job")
}
