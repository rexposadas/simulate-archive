package simulate

import (
	"testing"
	"fmt"
	"github.com/rexposadas/simulate"
)

func TestStats(t *testing.T){
	s := simulate.StatsObj{make(map[string]int)}
	s.Count =  map[string]int {"foo": 1, "buzz": 2}
	s.Add("foo") // add 1
	s.Add("foo") // add 1
	fmt.Printf("%+v\n", s.Count)
}