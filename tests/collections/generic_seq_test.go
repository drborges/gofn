package main

import (
	"github.com/drborges/gofn/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenericSeqForEach(t *testing.T) {
	seq := collections.NewGenericSeq(1, 2, 3)

	var items []int
	seq.ForEach(func(item collections.Any) {
		items = append(items, item.(int))
	})

	assert.Equal(t, items, []int{1, 2, 3})
}

func TestGenericSeqForEachResetsIterator(t *testing.T) {
	seq := collections.NewGenericSeq(1, 2, 3)

	var items []int
	seq.ForEach(func(item collections.Any) {
		items = append(items, item.(int))
	})

	seq.ForEach(func(item collections.Any) {
		items = append(items, item.(int)+3)
	})

	assert.Equal(t, items, []int{1, 2, 3, 4, 5, 6})
}

func TestGenericSeqMap(t *testing.T) {
	seq := collections.NewGenericSeq(1, 2, 3)

	newSeq := seq.Map(func(item collections.Any) collections.Any {
		return item.(int) + 1
	})

	assert.Equal(t, newSeq.AsArray(), []collections.Any{2, 3, 4})
}

func TestGenericSeqMapResetsIterator(t *testing.T) {
	seq := collections.NewGenericSeq(1, 2, 3)

	newSeq := seq.Map(func(item collections.Any) collections.Any {
		return item.(int) + 1
	})

	anotherSeq := seq.Map(func(item collections.Any) collections.Any {
		return item.(int) + 2
	})

	assert.Equal(t, newSeq.AsArray(), []collections.Any{2, 3, 4})
	assert.Equal(t, anotherSeq.AsArray(), []collections.Any{3, 4, 5})
}

func TestGenericSeqFindAll(t *testing.T) {
	seq := collections.NewGenericSeq(1, 2, 3, 4)

	evenNumbers := func(item collections.Any) bool {
		return item.(int)%2 == 0
	}

	result := seq.FindAll(evenNumbers)

	assert.Equal(t, result.AsArray(), []collections.Any{2, 4})
}

func TestGenericSeqFindAllResetsIterator(t *testing.T) {
	seq := collections.NewGenericSeq(1, 2, 3, 4)

	even := func(item collections.Any) bool {
		return item.(int)%2 == 0
	}

	odd := func(item collections.Any) bool {
		return item.(int)%2 != 0
	}

	evenNumbers := seq.FindAll(even)
	oddNumbers := seq.FindAll(odd)

	assert.Equal(t, evenNumbers.AsArray(), []collections.Any{2, 4})
	assert.Equal(t, oddNumbers.AsArray(), []collections.Any{1, 3})
}
