package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloats32(t *testing.T) {
	b := NewFloats32()

	b.Set("test1", float32(1))
	val, err := b.Get("test1")
	assert.NoError(t, err)
	assert.Equal(t, float32(1), val)

	val, err = b.Get("test2")
	assert.Equal(t, ErrNotFound, err)
	assert.Equal(t, float32(0), val)
}

func TestFloats64(t *testing.T) {
	b := NewFloats64()

	b.Set("test1", float64(1))
	val, err := b.Get("test1")
	assert.NoError(t, err)
	assert.Equal(t, float64(1), val)

	val, err = b.Get("test2")
	assert.Equal(t, ErrNotFound, err)
	assert.Equal(t, float64(0), val)
}
