package fn

import (
	"github.com/drborges/gofn/fn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackPush(t *testing.T) {
	var stack = fn.Stack{}

	stack.Push(1).Push(2)

	assert.Equal(t, fn.Stack{1, 2}, stack)
}

func TestStackPop(t *testing.T) {
	var stack = fn.Stack{1, 2, 3}

	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, fn.None(nil), stack.Pop())
}
