package main

// This is a sample application which shows a simple use case of
// simulate. This usecase hits the Google homepage once a second -
// that's the default interval behavior.
// The simulate defaults to writing to stdout/stderr.

import (
	"fmt"

	"github.com/rexposadas/simulate"
)


func main() {

	s := simulate.StatsObj{make(map[string]int)}
	s.Count =  map[string]int {"foo": 1, "buzz": 2}
	s.Add("foo") // add 1
	s.Add("foo") // add 1
	s.Sub("foo") // subtract 1
	fmt.Println(s.Count["foo"])
}
