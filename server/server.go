package server

import (
	"fmt"
	"time"
)

var jobs chan *Job

// Run runs the simulate server
func Run() {

	fmt.Println("server started")
	jobs = make(chan *Job, 1000)

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
				fmt.Printf("GET %s \n", j.URL)
				<-timer.C
			}
		}()
	}
}
