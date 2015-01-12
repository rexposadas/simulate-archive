package simulate

import (
	"sync"
)

type Counter struct {
	M    map[string]int
	lock sync.Mutex
}

func (c *Counter) Add(k string, v int) {

	c.lock.Lock()
	defer c.lock.Unlock()

	if val, ok := c.M[k]; ok {
		c.M[k] = val + v
	} else {
		c.M[k] = 1
	}
}

func NewCounter() *Counter {
	return &Counter{
		M: make(map[string]int),
	}
}
