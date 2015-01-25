package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeHasNext(t *testing.T) {
	rootValue := Any(1)
	tree := NewBinaryTree(NewBinaryTreeNode(&rootValue))

	assert.Equal(t, true, tree.HasNext())
	assert.Equal(t, false, EmptyBinaryTree.HasNext())
}

func TestBinaryTreeNext(t *testing.T) {
	tree := NewBinaryTree(
		NewBinaryTreeNodeWithChildren(1, NewBinaryTreeNode(2), NewBinaryTreeNode(3)))

	assert.Equal(t, 1, tree.Next().(*BinaryTreeNode).Value)
	assert.Equal(t, 2, tree.Next().(*BinaryTreeNode).Value)
	assert.Equal(t, 3, tree.Next().(*BinaryTreeNode).Value)
}

func TestDeepBinaryTreeNext(t *testing.T) {
	tree := NewBinaryTree(NewBinaryTreeNodeWithChildren(1,
		NewBinaryTreeNodeWithChildren(2, NewBinaryTreeNode(3), NewBinaryTreeNode(4)),
		NewBinaryTreeNodeWithChildren(5, NewBinaryTreeNode(6), NewBinaryTreeNode(7))))

	assert.Equal(t, 1, tree.Next().(*BinaryTreeNode).Value)
	assert.Equal(t, 2, tree.Next().(*BinaryTreeNode).Value)
	assert.Equal(t, 3, tree.Next().(*BinaryTreeNode).Value)
	assert.Equal(t, 4, tree.Next().(*BinaryTreeNode).Value)
	assert.Equal(t, 5, tree.Next().(*BinaryTreeNode).Value)
	assert.Equal(t, 6, tree.Next().(*BinaryTreeNode).Value)
	assert.Equal(t, 7, tree.Next().(*BinaryTreeNode).Value)
}

func TestBinaryTreeImplementsIterableInterface(t *testing.T) {
	tree := NewBinaryTree(NewBinaryTreeNode(1))

	assert.Implements(t, (*Iterable)(nil), tree)
}

func TestBinaryTreeIsTraversable(t *testing.T) {
	traversableTree := NewTraverser(NewBinaryTree(NewBinaryTreeNodeWithChildren(1,
		NewBinaryTreeNode(2),
		NewBinaryTreeNode(3))))

	sumAll := func(acc Any, next Any) Any {
		return acc.(int) + next.(*BinaryTreeNode).Value.(int)
	}

	assert.Equal(t, 6, traversableTree.Reduce(0)(sumAll))
}
