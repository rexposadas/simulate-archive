package simulate

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestAdd(t *testing.T) {

	s := NewCounter()

	s.Add("foo", 1)
	s.Add("foo", 1)
	s.Add("buzz", 1)

	assert.Equal(t, 2, s.Size("foo"))
	assert.Equal(t, 1, s.Size("buzz"))

	s.Add("buzz", 10)
	assert.Equal(t, 11, s.Size("buzz"))
}
