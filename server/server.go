package server

import (
	"fmt"
	"time"

	simhttp "github.com/rexposadas/simulate/http"
)

var jobs chan *Job

// Run runs the simulate server
func Run(port int) {

	fmt.Println("server started")
	jobs = make(chan *Job, 1000)

	go StartAPI(port)
	go consumer()
}

// Adds a GET job to the queue. This is the simplest job, where the
// server simply does a GET request to the supplied URL.
func Add(url string) {

	j := NewJob()
	j.URL = url

	jobs <- j
}

// Consumer gets a job and processes it.
// Each job received is a new routine.
func consumer() {

	for {
		j := <-jobs

		go func() {
			timer := time.NewTicker(j.Interval)

			for {
				resp, err := simhttp.Get(j.URL)
				if err != nil {
					fmt.Printf("got Error %+v on %s", err, j.URL)
				}
				fmt.Printf("GET '%s' - response time %f seconds. \n\n", j.URL, resp.Duration.Seconds())
				<-timer.C
			}
		}()

	}
}
