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

	assert.Equal(t, items, []int{1, 2, 3})
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

	assert.Equal(t, items, []int{1, 2, 3, 4, 5, 6})
}

func TestGenericSeqMap(t *testing.T) {
	seq := fn.NewGenericSeq(1, 2, 3)

	newSeq := seq.Map(func(item fn.Any) fn.Any {
		return item.(int) + 1
	})

	assert.Equal(t, newSeq.AsArray(), []fn.Any{2, 3, 4})
}

func TestGenericSeqMapResetsIterator(t *testing.T) {
	seq := fn.NewGenericSeq(1, 2, 3)

	newSeq := seq.Map(func(item fn.Any) fn.Any {
		return item.(int) + 1
	})

	anotherSeq := seq.Map(func(item fn.Any) fn.Any {
		return item.(int) + 2
	})

	assert.Equal(t, newSeq.AsArray(), []fn.Any{2, 3, 4})
	assert.Equal(t, anotherSeq.AsArray(), []fn.Any{3, 4, 5})
}

func TestGenericSeqFindAll(t *testing.T) {
	seq := fn.NewGenericSeq(1, 2, 3, 4)

	evenNumbers := func(item fn.Any) bool {
		return item.(int)%2 == 0
	}

	result := seq.FindAll(evenNumbers)

	assert.Equal(t, result.AsArray(), []fn.Any{2, 4})
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

	assert.Equal(t, evenNumbers.AsArray(), []fn.Any{2, 4})
	assert.Equal(t, oddNumbers.AsArray(), []fn.Any{1, 3})
}
