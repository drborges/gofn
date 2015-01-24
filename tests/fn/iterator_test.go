package main

import (
	"github.com/drborges/gofn/fn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator(t *testing.T) {
	items := []fn.Any{"a", "b", "c"}

	iterator := fn.NewIterator(items)

	assert.Equal(t, iterator.Next(), "a")
	assert.Equal(t, iterator.Next(), "b")
	assert.Equal(t, iterator.Next(), "c")
	assert.Equal(t, iterator.Next(), nil)
}

func TestIteratorWithDifferentInterfaceTypes(t *testing.T) {
	items := []fn.Any{"a", 1, true}

	iterator := fn.NewIterator(items)

	assert.Equal(t, iterator.Next(), "a")
	assert.Equal(t, iterator.Next(), 1)
	assert.Equal(t, iterator.Next(), true)
	assert.Equal(t, iterator.Next(), nil)
}

func TestIteratorDoesNotDestroyOriginalData(t *testing.T) {
	items := []fn.Any{"a", "b", "c"}

	iterator := fn.NewIterator(items)

	iterator.Next()
	iterator.Next()

	assert.Equal(t, []fn.Any{"a", "b", "c"}, items)
}
