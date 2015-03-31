package simulate

import (
	"fmt"
	"time"

	"github.com/franela/goreq"
)

// Config is the configuration struct
type Config struct {
	Type int
}

var (
	// Jobs is the job queue
	Jobs chan *Job

	// Port to run the server
	Port int

	// Metrics to keep various stats
	Metrics Stats
)

// NewConfig returns an initialized config
func NewConfig() Config {
	return Config{}
}

// Add adds (+1) to a key
func Add(k string) {
	Metrics.Tally(k, 1)
}

// Error sends an error to the metrics struct
func Error(err error, msg string) {
	Metrics.Error(err, msg)
}

// TrackResponse sends track messages to the metrics struct
func TrackResponse(req *goreq.Request, d time.Duration) {
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
