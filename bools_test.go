package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBools(t *testing.T) {
	b := NewBools()

	b.Set("test1", true)
	val, err := b.Get("test1")
	assert.NoError(t, err)
	assert.True(t, val)

	val, err = b.Get("test2")
	assert.Equal(t, ErrNotFound, err)
	assert.False(t, val)
}
