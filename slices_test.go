package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlices(t *testing.T) {
	b := NewStringSlices()

	src := []string{"val1"}

	b.Set("test1", src)
	val, err := b.Get("test1")
	assert.NoError(t, err)
	assert.Equal(t, src, val)
	assert.NotSame(t, src, val)

	val, err = b.Get("test2")
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, val)
}
