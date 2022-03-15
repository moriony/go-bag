package bag

import (
	"testing"

	"github.com/moriony/go-bag/unique"
	"github.com/stretchr/testify/assert"
)

func TestUniqueInts64(t *testing.T) {
	b := NewUniqueInts64()

	v := unique.NewInts64(1, 2, 3)
	b.Set("test1", v)
	val, err := b.Get("test1")
	assert.NoError(t, err)
	assert.Equal(t, v, val)

	val, err = b.Get("test2")
	assert.Equal(t, ErrNotFound, err)
	assert.IsType(t, unique.Ints64{}, val)

	b.Set("test2", v)
	val, err = b.Get("test2")
	assert.NoError(t, err)
	assert.Equal(t, v, val)
}
