package server

import (
	"fmt"
	"time"
)

var jobs chan string

// Run runs the simulate server
func Run() {

	fmt.Println("server started")
	jobs = make(chan string, 1000)

	go consumer()
}

// Add adds a job to the queue
func Add(s string) {

	timer := time.NewTicker(time.Second)

	go func() {
		for {
			jobs <- s
			<-timer.C
		}
	}()
}

// Consumer gets a job and processes it.
func consumer() {
	for {
		j := <-jobs

		fmt.Printf("processing job %s \n", j)
	}
}
