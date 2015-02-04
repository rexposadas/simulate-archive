package simulate

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Type int
}

var (
	Jobs    chan *Job
	Port    int
	Metrics Stats
)

func NewConfig() Config {
	return Config{}
}

func Add(k string) {
	Metrics.Tally(k, 1)
}

func Error(err error, msg string) {
	Metrics.Error(err, msg)
}

func TrackResponse(req *http.Request, d time.Duration) {
	Metrics.TrackResponse(req, d)
}

// Run runs the simulate server
func Run(c Config) {
	Jobs = make(chan *Job, 1000)

	// todo: set via command line argument
	if c.Type == INFLUXDB {
		Metrics = NewInfluxDB()
	} else {
		Metrics = NewPrintStats() // default
	}

	go Metrics.Run()

	fmt.Println("Simulator started \n\n")

	go worker()
}

// Worker runs the jobs.
// Each job received is a new routine.
func worker() {

	for {
		j := <-Jobs
		go j.Actor.Run()
	}
}
