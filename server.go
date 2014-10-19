package simulate

import (
	"fmt"
	"time"
)

var (
	Jobs chan *Job
	Port int
)

// Run runs the simulate server
func Run() {
	Jobs = make(chan *Job, 1000)

	fmt.Println("Simulator started \n\n")

	// The API is optional
	// todo: use CL flag
	if Port > 0 {
		go StartAPI(Port)
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
