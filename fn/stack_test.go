package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackPush(t *testing.T) {
	var stack = Stack{}

	stack.Push(1).Push(2)

	assert.Equal(t, Stack{1, 2}, stack)
}

func TestStackPop(t *testing.T) {
	var stack = Stack{1, 2, 3}

	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, None(nil), stack.Pop())
}

func TestStackLength(t *testing.T) {
	stack := Stack{1, 2}

	assert.Equal(t, 2, stack.Length())
	assert.Equal(t, 0, EmptyStack.Length())
}
