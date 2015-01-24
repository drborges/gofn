package fn

type BinaryTreeNode struct {
	Value  *Any
	LChild *BinaryTreeNode
	RChild *BinaryTreeNode
}

var EmptyNode = &BinaryTreeNode{}

func NewBinaryTreeNode(value *Any) *BinaryTreeNode {
	return &BinaryTreeNode{
		Value:  value,
		LChild: EmptyNode,
		RChild: EmptyNode,
	}
}

func NewBinaryTreeNodeWithChildren(value *Any, children ...*BinaryTreeNode) *BinaryTreeNode {
	lChild, rChild := EmptyNode, EmptyNode
	if len(children) == 1 {
		lChild = children[0]
	}

	if len(children) == 2 {
		lChild = children[0]
		rChild = children[1]
	}

	return &BinaryTreeNode{
		Value:  value,
		LChild: lChild,
		RChild: rChild,
	}
}

type BinaryTree struct {
	Root   *BinaryTreeNode
	Fringe *Stack
}

var EmptyBinaryTree = &BinaryTree{Fringe: EmptyStack}

func NewBinaryTree(root *BinaryTreeNode) *BinaryTree {
	return &BinaryTree{
		Root:   root,
		Fringe: &Stack{root},
	}
}

func (t *BinaryTree) HasNext() bool {
	return t.Fringe.Length() > 0
}

func (t *BinaryTree) Next() Any {
	if !t.HasNext() {
		return EmptyNode
	}

	node := t.Fringe.Pop().(*BinaryTreeNode)

	if node.RChild != EmptyNode {
		t.Fringe.Push(node.RChild)
	}

	if node.LChild != EmptyNode {
		t.Fringe.Push(node.LChild)
	}

	return node
}
