package simulate

import (
	"sync"
)

type Counter struct {
	M map[string]int
	sync.Mutex
}

func (c *Counter) Add(k string, v int) {

	c.Lock()
	defer c.Unlock()

	if val, ok := c.M[k]; ok {
		c.M[k] = val + v
	} else {
		c.M[k] = 1
	}
}

func (c *Counter) Size(k string) int {
	c.Lock()
	defer c.Unlock()

	if val, ok := c.M[k]; ok {
		return val
	}

	return 0
}

func NewCounter() *Counter {
	return &Counter{
		M: make(map[string]int),
	}
}
