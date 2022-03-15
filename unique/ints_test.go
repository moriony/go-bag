package unique

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts64(t *testing.T) {
	b := NewInts64(1, 2, 3, 3, 3, 3)

	assert.Equal(t, b.Len(), 3)
	assert.Len(t, b.Val(), 3)
	assert.True(t, b.Contains(1))
	assert.True(t, b.Contains(2))
	assert.True(t, b.Contains(3))
	assert.False(t, b.Contains(4))
}
