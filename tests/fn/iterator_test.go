package main

import (
	"github.com/drborges/gofn/fn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator(t *testing.T) {
	iterator := fn.NewIterator("a", "b", "c")

	assert.Equal(t, iterator.Next(), "a")
	assert.Equal(t, iterator.Next(), "b")
	assert.Equal(t, iterator.Next(), "c")
	assert.Equal(t, iterator.Next(), nil)
}

func TestIteratorWithDifferentInterfaceTypes(t *testing.T) {
	iterator := fn.NewIterator("a", 1, true)

	assert.Equal(t, iterator.Next(), "a")
	assert.Equal(t, iterator.Next(), 1)
	assert.Equal(t, iterator.Next(), true)
	assert.Equal(t, iterator.Next(), nil)
}

func TestIteratorReset(t *testing.T) {
	iterator := fn.NewIterator("a", 1, true)

	assert.Equal(t, iterator.Next(), "a")
	assert.Equal(t, iterator.Next(), 1)

	iterator.Reset()

	assert.Equal(t, iterator.Next(), "a")
}

func TestIteratorLength(t *testing.T) {
	iterator := fn.NewIterator(1, 2)

	assert.Equal(t, iterator.Length(), 2)
}

func TestIteratorToArray(t *testing.T) {
	iterator := fn.NewIterator(1, 2, 3)

	assert.Equal(t, iterator.AsArray(), []fn.Any{1, 2, 3})
}
