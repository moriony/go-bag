package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrings(t *testing.T) {
	b := NewStrings()

	b.Set("test1", "val1")
	val, err := b.Get("test1")
	assert.NoError(t, err)
	assert.Equal(t, "val1", val)

	val, err = b.Get("test2")
	assert.Equal(t, ErrNotFound, err)
	assert.Equal(t, "", val)
}
