package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator(t *testing.T) {
	items := []Any{"a", "b", "c"}

	iterator := NewIterator(items)

	assert.Equal(t, iterator.Next(), "a")
	assert.Equal(t, iterator.Next(), "b")
	assert.Equal(t, iterator.Next(), "c")
	assert.Equal(t, iterator.Next(), nil)
}

func TestIteratorWithDifferentInterfaceTypes(t *testing.T) {
	items := []Any{"a", 1, true}

	iterator := NewIterator(items)

	assert.Equal(t, iterator.Next(), "a")
	assert.Equal(t, iterator.Next(), 1)
	assert.Equal(t, iterator.Next(), true)
	assert.Equal(t, iterator.Next(), nil)
}

func TestIteratorDoesNotDestroyOriginalData(t *testing.T) {
	items := []Any{"a", "b", "c"}

	iterator := NewIterator(items)

	iterator.Next()
	iterator.Next()

	assert.Equal(t, []Any{"a", "b", "c"}, items)
}
