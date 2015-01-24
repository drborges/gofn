package fn

import (
	"github.com/drborges/gofn/fn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	var stack = fn.Stack{1, 2, 3}
	var firstItem, secondItem, thirdItem fn.Any

	stack = stack.Pop(&firstItem).Pop(&secondItem).Pop(&thirdItem)

	assert.Equal(t, []fn.Any{}, stack)
	assert.Equal(t, 3, firstItem)
	assert.Equal(t, 2, secondItem)
	assert.Equal(t, 1, thirdItem)
}
