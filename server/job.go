package server

import (
	"time"
)

const (
	METHOD_GET = "GET"
)

type call func() error

type Job struct {
	Interval time.Duration
	Run      call
}

func NewJob() *Job {
	return &Job{
		Interval: time.Second * 2,
	}
}
