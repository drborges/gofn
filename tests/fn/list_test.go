package main

import (
	"github.com/drborges/gofn/fn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListForEach(t *testing.T) {
	list := fn.NewList(1, 2, 3)

	var items []int
	list.ForEach(func(item fn.Any) {
		items = append(items, item.(int))
	})

	assert.Equal(t, []int{1, 2, 3}, items)
}

func TestListMap(t *testing.T) {
	list := fn.NewList(1, 2, 3)

	newList := list.Map(func(item fn.Any) fn.Any {
		return item.(int) + 1
	})

	assert.Equal(t, []fn.Any{2, 3, 4}, newList)
}

func TestListFind(t *testing.T) {
	list := fn.NewList(1, 2, 3, 4)

	firstEven := func(item fn.Any) bool {
		return item.(int)%2 == 0
	}

	assert.Equal(t, 2, list.Find(firstEven))
}

func TestListFilter(t *testing.T) {
	list := fn.NewList(1, 2, 3, 4)

	evenNumbers := func(item fn.Any) bool {
		return item.(int)%2 == 0
	}

	assert.Equal(t, []fn.Any{2, 4}, list.Filter(evenNumbers))
}

func TestListReduce(t *testing.T) {
	list := fn.NewList(1, 2, 3)

	total := list.Reduce(0)(func(sum, next fn.Any) fn.Any {
		return sum.(int) + next.(int)
	})

	assert.Equal(t, 6, total)
}

func TestListFlatten(t *testing.T) {
	list := fn.NewList([]fn.Any{1, 2}, 3)

	assert.Equal(t, []fn.Any{1, 2, 3}, list.Flatten())
}

func TestListAppend(t *testing.T) {
	list := fn.NewList(1, 2)

	assert.Equal(t, []fn.Any{1, 2, 3}, list.Append(3))
}

func TestListHead(t *testing.T) {
	list := fn.NewList(1, 2, 3)

	assert.Equal(t, 1, list.Head())
}

func TestListTail(t *testing.T) {
	list := fn.NewList(1, 2, 3)

	assert.Equal(t, []fn.Any{2, 3}, list.Tail())
}

func TestListIter(t *testing.T) {
	it := fn.NewList(1, 2, 3).Iter()

	assert.Equal(t, 1, it.Next())
	assert.Equal(t, 2, it.Next())
	assert.Equal(t, 3, it.Next())
	assert.Equal(t, nil, it.Next())
}
