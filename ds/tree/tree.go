package tree

import "galaxyzeta.com/algo/compare"

const (
	left  = -1
	right = 1
)

// TNode represents the node of a binary tree.
type TNode struct {
	lchild *TNode
	rchild *TNode
	key    interface{}
	val    interface{}
}

// LChild returns the left child of a tree.
func (t *TNode) LChild() *TNode {
	return t.lchild
}

// RChild returns the left child of a tree.
func (t *TNode) RChild() *TNode {
	return t.lchild
}

// SetLeftLeaf adds a leaf node to the left side of a given tree node.
func (t *TNode) SetLeftLeaf(key interface{}, val interface{}) {
	t.lchild = newNode(key, val)
}

// SetRightLeaf adds a leaf node to the right side of a given tree node.
func (t *TNode) SetRightLeaf(key interface{}, val interface{}) {
	t.rchild = newNode(key, val)
}

// SetLChild sets the left child node.
func (t *TNode) SetLChild(n *TNode) {
	t.lchild = n
}

// SetRChild sets the right child node.
func (t *TNode) SetRChild(n *TNode) {
	t.rchild = n
}

// newNode returns a new TreeNode.
func newNode(key interface{}, val interface{}) *TNode {
	return &TNode{key: key, val: val}
}

// free sets a TreeNode to blank.
func (t *TNode) free() {
	t.val = nil
	t.key = nil
	t.lchild = nil
	t.rchild = nil
}

// keySearch searches expected key from the start TreeNode, using given comparison function.
//
// Return triple values:
//
// - Previous TNode before expected Node.
//
// - Expected Node.
//
// - Whether the expected node is the left child of the previous node
func keySearch(startFrom *TNode, key interface{}, cmpFunc compare.CompareFunc) (prev *TNode, current *TNode, isleft bool) {
	current = startFrom
	if current == nil {
		return nil, nil, false
	}
	flag := true
	prev = nil
	for current != nil {
		res := cmpFunc(current.key, key)
		switch res {
		case compare.Less: // val > current.val
			prev = current
			current = current.rchild
			flag = false
		case compare.Greater: // val <= current.val
			prev = current
			current = current.lchild
			flag = true
		case compare.Equal:
			// modify value directly.
			return prev, current, flag
		}
	}
	return nil, nil, false
}
