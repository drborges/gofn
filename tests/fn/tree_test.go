package fn

import (
	"github.com/drborges/gofn/fn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeHasNext(t *testing.T) {
	rootValue := fn.Any(1)
	tree := fn.NewBinaryTree(fn.NewBinaryTreeNode(&rootValue))

	assert.Equal(t, true, tree.HasNext())
	assert.Equal(t, false, fn.EmptyBinaryTree.HasNext())
}

func TestBinaryTreeNext(t *testing.T) {
	rootValue, lChildValue, rChildValue := fn.Any(1), fn.Any(2), fn.Any(3)
	lChild := fn.NewBinaryTreeNode(&lChildValue)
	rChild := fn.NewBinaryTreeNode(&rChildValue)
	root := fn.NewBinaryTreeNodeWithChildren(&rootValue, lChild, rChild)

	tree := fn.NewBinaryTree(root)

	assert.Equal(t, root, tree.Next())
	assert.Equal(t, lChild, tree.Next())
	assert.Equal(t, rChild, tree.Next())
}
