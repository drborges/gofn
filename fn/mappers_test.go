package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIter2Array(t *testing.T) {
	iter := NewIteratorFromItems(1, 2, 3)

	assert.Equal(t, []Any{1, 2, 3}, Iter2Array(iter))
}
