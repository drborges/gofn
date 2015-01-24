package fn

import "testing"
import (
	"github.com/drborges/gofn/fn"
	"github.com/stretchr/testify/assert"
)

func TestIter2Array(t *testing.T) {
	iter := fn.NewIteratorFromItems(1, 2, 3)

	assert.Equal(t, []fn.Any{1, 2, 3}, fn.Iter2Array(iter))
}
