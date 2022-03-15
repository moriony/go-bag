package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnique(t *testing.T) {
	u := NewUnique()
	assert.NotNil(t, u.Ints64())
}
