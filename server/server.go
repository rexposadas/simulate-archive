package server

import (
	"fmt"
	"time"
)

var (
	Jobs chan *Job
)

// Run runs the simulate server
func Run(port int) {
	Jobs = make(chan *Job, 1000)

	fmt.Println("Simulator started \n\n")

	// The API is optional
	// todo: use CL flag
	if port > 0 {
		go StartAPI(port)
	}

	go consumer()
}

// Adds a GET job to the queue. This is the simplest job, where the
// server simply does a GET request to the supplied URL.
func Add(url string) {

	j := NewJob()
	Jobs <- j
}

// Consumer gets a job and processes it.
// Each job received is a new routine.
func consumer() {

	for {
		j := <-Jobs

		go func() {
			timer := time.NewTicker(j.Interval)

			for {
				j.Run()
				<-timer.C
			}
		}()

	}
}
