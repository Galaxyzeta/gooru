package tree

import (
	"galaxyzeta.com/algo/compare"
	"galaxyzeta.com/ds/list"
	"galaxyzeta.com/util/alias"
)

const (
	left  = -1
	right = 1
)

// ElemType represents any type.
type ElemType interface{}

// TNode represents the node of a binary tree.
type TNode struct {
	lchild *TNode
	rchild *TNode
	val    ElemType
}

// BinaryTree represents a two-branched tree data structure.
type BinaryTree struct {
	tree    *TNode
	cmpFunc compare.CompareFunc
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
func (t *TNode) SetLeftLeaf(val interface{}) {
	t.lchild = newNode(val)
}

// SetRightLeaf adds a leaf node to the right side of a given tree node.
func (t *TNode) SetRightLeaf(val interface{}) {
	t.rchild = newNode(val)
}

// SetLChild sets the left child node.
func (t *TNode) SetLChild(n *TNode) {
	t.lchild = n
}

// SetRChild sets the right child node.
func (t *TNode) SetRChild(n *TNode) {
	t.rchild = n
}

// NewBST builds a BST using given arr.
func NewBST(lambda compare.CompareFunc) *BinaryTree {
	if lambda == nil {
		return &BinaryTree{tree: nil, cmpFunc: compare.BasicCompare}
	}
	return &BinaryTree{tree: nil, cmpFunc: lambda}
}

// Insert into a BST using given comparator
func (bst *BinaryTree) Insert(val ElemType) {
	if bst.tree == nil {
		// initial insertion
		bst.tree = newNode(val)
	} else {
		// binary search descend
		current := bst.tree
		prev := current
		flag := left
		for current != nil {
			res := bst.cmpFunc(current.val, val)
			switch res {
			case compare.Greater:
				prev = current
				current = current.rchild
				flag = left
			default:
				prev = current
				current = current.lchild
				flag = left
			}
		}
		if flag == left {
			prev.lchild = newNode(val)
		} else if flag == right {
			prev.rchild = newNode(val)
		} else {
			panic("unreachable situation.")
		}
	}
}

// PreOrderTraverse the given binary tree with a consumer.
func (bst *BinaryTree) PreOrderTraverse(fn alias.P1Consumer) {
	bst.doPreOrderTraverse(bst.tree, fn)
}

// InOrderTraverse the given binary tree with a consumer.
func (bst *BinaryTree) InOrderTraverse(fn alias.P1Consumer) {
	bst.doInOrderTraverse(bst.tree, fn)
}

// PostOrderTraverse the given binary tree with a consumer.
func (bst *BinaryTree) PostOrderTraverse(fn alias.P1Consumer) {
	bst.doPostOrderTraverse(bst.tree, fn)
}

// LevelOrderTraverse the given binary tree with a consumer.
func (bst *BinaryTree) LevelOrderTraverse(fn alias.P1Consumer) {
	var queue list.Queue = list.NewSinglyLinkedList()
	if bst.tree == nil {
		return
	}
	queue.Offer(bst.tree)
	for !queue.IsEmpty() {
		for i := queue.Size(); i > 0; i-- {
			elem := queue.Poll().(*TNode)
			fn(elem)
			if elem.lchild != nil {
				queue.Offer(elem.lchild)
			}
			if elem.rchild != nil {
				queue.Offer(elem.rchild)
			}
		}
	}
}

func (bst *BinaryTree) doPreOrderTraverse(node *TNode, fn alias.P1Consumer) {
	if node == nil {
		return
	}
	fn(node.val)
	bst.doPreOrderTraverse(node.lchild, fn)
	bst.doPreOrderTraverse(node.rchild, fn)
}

func (bst *BinaryTree) doInOrderTraverse(node *TNode, fn alias.P1Consumer) {
	if node == nil {
		return
	}
	bst.doPreOrderTraverse(node.lchild, fn)
	fn(node.val)
	bst.doPreOrderTraverse(node.rchild, fn)
}

func (bst *BinaryTree) doPostOrderTraverse(node *TNode, fn alias.P1Consumer) {
	if node == nil {
		return
	}
	bst.doPreOrderTraverse(node.lchild, fn)
	bst.doPreOrderTraverse(node.rchild, fn)
	fn(node.val)
}

func newNode(val ElemType) *TNode {
	return &TNode{val: val}
}
