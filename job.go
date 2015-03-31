package simulate

import (
	"time"
)

type call func() error

// ActorImp is an interface which indicates an actor that
// the simulator runs.
type ActorImp interface {
	Run() error
}

// Job is a job which simulate runs.  Jobs are sent to the Jobs channel
// for processing.
type Job struct {
	Iteration int
	Delay     time.Duration
	Actor     ActorImp
}

// NewJob returns a job with a default values.
func NewJob() *Job {
	return &Job{
		Iteration: 0,
		Delay:     time.Second * 2, // arbitrary interval to do the job.
	}
}
