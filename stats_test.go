package simulate

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestAddSub(t *testing.T) {

	s := New()

	s.Add("foo")
	s.Add("foo")
	s.Add("buzz")

	assert.Equal(t, 2, s.Count["foo"])
	assert.Equal(t, 1, s.Count["buzz"])

	s.SimpleMath("buzz", 10)
	assert.Equal(t, 11, s.Count["buzz"])
}
