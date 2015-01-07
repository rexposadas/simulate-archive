package simulate

import (
	"sync"
)

// StatsObject is in charge of keep track of statistics
type StatsObj struct {
	Counter
}
type Counter struct {

	// this variable keeps a map of strings and it's count.
	Count map[string]int
	sync.RWMutex
}

func New() *StatsObj {
	s := &StatsObj{}
	s.Count = make(map[string]int)
	return s
}

// Tick adds a one (+1) to the Map for a given key
func (c *Counter) Add(t string) {
	c.SimpleMath(t, 1)
}

func (c *Counter) Sub(t string) {
	c.SimpleMath(t, -1)
}

// simpleMath adds/substracts an arbitrary amount from a map key
func (c *Counter) SimpleMath(t string, count int) {
	c.Lock()
	defer c.Unlock()

	if val, ok := c.Count[t]; ok {
		c.Count[t] = val + count
	} else {
		c.Count[t] = count
	}
}
