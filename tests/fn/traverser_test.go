package main

import (
	"github.com/drborges/gofn/fn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTraverserForEach(t *testing.T) {
	traverser := fn.NewTraverser(fn.NewIteratorFromItems(1, 2, 3))

	var items []int
	traverser.ForEach(func(item fn.Any) {
		items = append(items, item.(int))
	})

	assert.Equal(t, []int{1, 2, 3}, items)
}

func TestTraverserMap(t *testing.T) {
	traverser := fn.NewTraverser(fn.NewIteratorFromItems(1, 2, 3, 4))

	newTraverser := traverser.Map(func(item fn.Any) fn.Any {
		return item.(int) + 1
	})

	assert.Equal(t, []fn.Any{2, 3, 4, 5}, fn.Iter2Array(newTraverser))
}

func TestTraverserFind(t *testing.T) {
	traverser := fn.NewTraverser(fn.NewIteratorFromItems(1, 2, 3, 4))

	firstEven := func(item fn.Any) bool {
		return item.(int)%2 == 0
	}

	assert.Equal(t, 2, traverser.Find(firstEven))
}

func TestTraverserFilter(t *testing.T) {
	traverser := fn.NewTraverser(fn.NewIteratorFromItems(1, 2, 3, 4))

	evenNumbers := func(item fn.Any) bool {
		return item.(int)%2 == 0
	}

	assert.Equal(t, []fn.Any{2, 4}, fn.Iter2Array(traverser.Filter(evenNumbers)))
}

func TestTraverserReduce(t *testing.T) {
	traverser := fn.NewTraverser(fn.NewIteratorFromItems(1, 2, 3))

	total := traverser.Reduce(0)(func(sum, next fn.Any) fn.Any {
		return sum.(int) + next.(int)
	})

	assert.Equal(t, 6, total)
}

func TestTraverserFlatten(t *testing.T) {
	traverser := fn.NewTraverser(fn.NewIteratorFromItems([]fn.Any{1, 2}, 3))

	assert.Equal(t, []fn.Any{1, 2, 3}, fn.Iter2Array(traverser.Flatten()))
}
