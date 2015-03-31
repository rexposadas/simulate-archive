package simulate

import (
	"sync"
)

// Counter counts various metrics
type Counter struct {
	M map[string]int
	sync.Mutex
}

// Add adds a tally to a key
func (c *Counter) Add(k string, v int) {

	c.Lock()
	defer c.Unlock()

	if val, ok := c.M[k]; ok {
		c.M[k] = val + v
	} else {
		c.M[k] = 1
	}
}

// Size get the count on key
func (c *Counter) Size(k string) int {
	c.Lock()
	defer c.Unlock()

	if val, ok := c.M[k]; ok {
		return val
	}

	return 0
}

// NewCounter returns an initialized counter
func NewCounter() *Counter {
	return &Counter{
		M: make(map[string]int),
	}
}
