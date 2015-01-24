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
	tree := fn.NewBinaryTree(
		fn.NewBinaryTreeNodeWithChildren(1, fn.NewBinaryTreeNode(2), fn.NewBinaryTreeNode(3)))

	assert.Equal(t, 1, tree.Next().(*fn.BinaryTreeNode).Value)
	assert.Equal(t, 2, tree.Next().(*fn.BinaryTreeNode).Value)
	assert.Equal(t, 3, tree.Next().(*fn.BinaryTreeNode).Value)
}

func TestDeepBinaryTreeNext(t *testing.T) {
	tree := fn.NewBinaryTree(fn.NewBinaryTreeNodeWithChildren(1,
		fn.NewBinaryTreeNodeWithChildren(2, fn.NewBinaryTreeNode(3), fn.NewBinaryTreeNode(4)),
		fn.NewBinaryTreeNodeWithChildren(5, fn.NewBinaryTreeNode(6), fn.NewBinaryTreeNode(7))))

	assert.Equal(t, 1, tree.Next().(*fn.BinaryTreeNode).Value)
	assert.Equal(t, 2, tree.Next().(*fn.BinaryTreeNode).Value)
	assert.Equal(t, 3, tree.Next().(*fn.BinaryTreeNode).Value)
	assert.Equal(t, 4, tree.Next().(*fn.BinaryTreeNode).Value)
	assert.Equal(t, 5, tree.Next().(*fn.BinaryTreeNode).Value)
	assert.Equal(t, 6, tree.Next().(*fn.BinaryTreeNode).Value)
	assert.Equal(t, 7, tree.Next().(*fn.BinaryTreeNode).Value)
}

func TestBinaryTreeImplementsIterableInterface(t *testing.T) {
	tree := fn.NewBinaryTree(fn.NewBinaryTreeNode(1))

	assert.Implements(t, (*fn.Iterable)(nil), tree)
}

func TestBinaryTreeIsTraversable(t *testing.T) {
	traversableTree := fn.NewTraverser(fn.NewBinaryTree(fn.NewBinaryTreeNodeWithChildren(1,
		fn.NewBinaryTreeNode(2),
		fn.NewBinaryTreeNode(3))))

	sumAll := func(acc fn.Any, next fn.Any) fn.Any {
		return acc.(int) + next.(*fn.BinaryTreeNode).Value.(int)
	}

	assert.Equal(t, 6, traversableTree.Reduce(0)(sumAll))
}
