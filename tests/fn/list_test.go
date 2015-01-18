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

func TestListForEachResetsIterator(t *testing.T) {
	list := fn.NewList(1, 2, 3)

	var items []int
	list.ForEach(func(item fn.Any) {
		items = append(items, item.(int))
	})

	list.ForEach(func(item fn.Any) {
		items = append(items, item.(int)+3)
	})

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, items)
}

func TestListMap(t *testing.T) {
	list := fn.NewList(1, 2, 3)

	newList := list.Map(func(item fn.Any) fn.Any {
		return item.(int) + 1
	})

	assert.Equal(t, []fn.Any{2, 3, 4}, newList.AsArray())
}

func TestListMapResetsIterator(t *testing.T) {
	list := fn.NewList(1, 2, 3)

	newList := list.Map(func(item fn.Any) fn.Any {
		return item.(int) + 1
	})

	anotherList := list.Map(func(item fn.Any) fn.Any {
		return item.(int) + 2
	})

	assert.Equal(t, []fn.Any{2, 3, 4}, newList.AsArray())
	assert.Equal(t, []fn.Any{3, 4, 5}, anotherList.AsArray())
}

func TestListFind(t *testing.T) {
	list := fn.NewList(1, 2, 3, 4)

	firstEven := func(item fn.Any) bool {
		return item.(int)%2 == 0
	}

	assert.Equal(t, 2, list.Find(firstEven))
}

func TestListFindResetsIterator(t *testing.T) {
	list := fn.NewList(1, 2, 3, 4)

	firstEven := func(item fn.Any) bool {
		return item.(int)%2 == 0
	}

	firstOdd := func(item fn.Any) bool {
		return item.(int)%2 != 0
	}

	assert.Equal(t, 2, list.Find(firstEven))
	assert.Equal(t, 3, list.Find(firstOdd))
}

func TestListFilter(t *testing.T) {
	list := fn.NewList(1, 2, 3, 4)

	evenNumbers := func(item fn.Any) bool {
		return item.(int)%2 == 0
	}

	assert.Equal(t, []fn.Any{2, 4}, list.Filter(evenNumbers).AsArray())
}

func TestListFilterResetsIterator(t *testing.T) {
	list := fn.NewList(1, 2, 3, 4)

	even := func(item fn.Any) bool {
		return item.(int)%2 == 0
	}

	odd := func(item fn.Any) bool {
		return item.(int)%2 != 0
	}

	assert.Equal(t, []fn.Any{2, 4}, list.Filter(even).AsArray())
	assert.Equal(t, []fn.Any{1, 3}, list.Filter(odd).AsArray())
}

func TestListReduce(t *testing.T) {
	list := fn.NewList(1, 2, 3)

	total := list.Reduce(0)(func(sum, next fn.Any) fn.Any {
		return sum.(int) + next.(int)
	})

	assert.Equal(t, 6, total)
}

func TestListReduceResetsIterator(t *testing.T) {
	list := fn.NewList(1, 2, 3)

	total1 := list.Reduce(0)(func(sum, next fn.Any) fn.Any {
		return sum.(int) + next.(int)
	})

	total2 := list.Reduce(0)(func(sum, next fn.Any) fn.Any {
		return sum.(int) + next.(int)
	})

	assert.Equal(t, 12, total1.(int)+total2.(int))
}
