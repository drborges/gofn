package fn

import (
	"github.com/drborges/gofn/fn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeHasNext(t *testing.T) {
	tree := fn.NewBinaryTree(fn.NewBinaryTreeNode(&fn.Any(1)))

	assert.Equal(t, true, tree.HasNext())
	assert.Equal(t, false, fn.EmptyBinaryTree.HasNext())
}
