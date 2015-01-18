package main

import (
	"github.com/drborges/gofn/fn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenericIterator(t *testing.T) {
	iterator := fn.NewGenericIterable("a", "b", "c")

	assert.Equal(t, iterator.Next(), "a")
	assert.Equal(t, iterator.Next(), "b")
	assert.Equal(t, iterator.Next(), "c")
	assert.Equal(t, iterator.Next(), nil)
}

func TestGenericIteratorWithDifferentInterfaceTypes(t *testing.T) {
	iterator := fn.NewGenericIterable("a", 1, true)

	assert.Equal(t, iterator.Next(), "a")
	assert.Equal(t, iterator.Next(), 1)
	assert.Equal(t, iterator.Next(), true)
	assert.Equal(t, iterator.Next(), nil)
}

func TestGenericIteratorReset(t *testing.T) {
	iterator := fn.NewGenericIterable("a", 1, true)

	assert.Equal(t, iterator.Next(), "a")
	assert.Equal(t, iterator.Next(), 1)

	iterator.Reset()

	assert.Equal(t, iterator.Next(), "a")
}

func TestGenericIteratorLength(t *testing.T) {
	iterator := fn.NewGenericIterable(1, 2)

	assert.Equal(t, iterator.Length(), 2)
}

func TestGenericIteratorToArray(t *testing.T) {
	iterator := fn.NewGenericIterable(1, 2, 3)

	assert.Equal(t, iterator.AsArray(), []fn.Any{1, 2, 3})
}
