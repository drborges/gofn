package main

import (
	"github.com/drborges/gofn/fn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenericSeqForEach(t *testing.T) {
	seq := fn.NewGenericSeq(1, 2, 3)

	var items []int
	seq.ForEach(func(item fn.Any) {
		items = append(items, item.(int))
	})

	assert.Equal(t, []int{1, 2, 3}, items)
}

func TestGenericSeqForEachResetsIterator(t *testing.T) {
	seq := fn.NewGenericSeq(1, 2, 3)

	var items []int
	seq.ForEach(func(item fn.Any) {
		items = append(items, item.(int))
	})

	seq.ForEach(func(item fn.Any) {
		items = append(items, item.(int)+3)
	})

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, items)
}

func TestGenericSeqMap(t *testing.T) {
	seq := fn.NewGenericSeq(1, 2, 3)

	newSeq := seq.Map(func(item fn.Any) fn.Any {
		return item.(int) + 1
	})

	assert.Equal(t, []fn.Any{2, 3, 4}, newSeq.AsArray())
}

func TestGenericSeqMapResetsIterator(t *testing.T) {
	seq := fn.NewGenericSeq(1, 2, 3)

	newSeq := seq.Map(func(item fn.Any) fn.Any {
		return item.(int) + 1
	})

	anotherSeq := seq.Map(func(item fn.Any) fn.Any {
		return item.(int) + 2
	})

	assert.Equal(t, []fn.Any{2, 3, 4}, newSeq.AsArray())
	assert.Equal(t, []fn.Any{3, 4, 5}, anotherSeq.AsArray())
}

func TestGenericSeqFindAll(t *testing.T) {
	seq := fn.NewGenericSeq(1, 2, 3, 4)

	evenNumbers := func(item fn.Any) bool {
		return item.(int)%2 == 0
	}

	result := seq.FindAll(evenNumbers)

	assert.Equal(t, []fn.Any{2, 4}, result.AsArray())
}

func TestGenericSeqFindAllResetsIterator(t *testing.T) {
	seq := fn.NewGenericSeq(1, 2, 3, 4)

	even := func(item fn.Any) bool {
		return item.(int)%2 == 0
	}

	odd := func(item fn.Any) bool {
		return item.(int)%2 != 0
	}

	evenNumbers := seq.FindAll(even)
	oddNumbers := seq.FindAll(odd)

	assert.Equal(t, []fn.Any{2, 4}, evenNumbers.AsArray())
	assert.Equal(t, []fn.Any{1, 3}, oddNumbers.AsArray())
}

func TestGenericSeqReduce(t *testing.T) {
	seq := fn.NewGenericSeq(1, 2, 3)

	total := seq.Reduce(0)(func(sum, next fn.Any) fn.Any {
		return sum.(int) + next.(int)
	})

	assert.Equal(t, 6, total)
}

func TestGenericSeqReduceResetsIterator(t *testing.T) {
	seq := fn.NewGenericSeq(1, 2, 3)

	total1 := seq.Reduce(0)(func(sum, next fn.Any) fn.Any {
		return sum.(int) + next.(int)
	})

	total2 := seq.Reduce(0)(func(sum, next fn.Any) fn.Any {
		return sum.(int) + next.(int)
	})

	assert.Equal(t, 12, total1.(int)+total2.(int))
}
