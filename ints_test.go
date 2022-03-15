package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {
	b := NewInts()

	b.Set("test1", 1)
	val, err := b.Get("test1")
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = b.Get("test2")
	assert.Equal(t, ErrNotFound, err)
	assert.Equal(t, 0, val)
}

func TestInts64(t *testing.T) {
	b := NewInts64()

	b.Set("test1", int64(1))
	val, err := b.Get("test1")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)

	val, err = b.Get("test2")
	assert.Equal(t, ErrNotFound, err)
	assert.Equal(t, int64(0), val)
}
