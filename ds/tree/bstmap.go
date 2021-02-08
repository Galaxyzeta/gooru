package tree

import (
	"galaxyzeta.com/algo/compare"
	"galaxyzeta.com/ds/list"
	"galaxyzeta.com/util/alias"
	"galaxyzeta.com/util/common"
)

// BSTMap represents a two-branched tree data structure.
type BSTMap struct {
	tree    *TNode
	cmpFunc compare.CompareFunc
	size    int
}

// NewBST creates a new BST.
func NewBST(lambda compare.CompareFunc) *BSTMap {
	if lambda == nil {
		return &BSTMap{tree: nil, cmpFunc: compare.BasicCompare}
	}
	return &BSTMap{tree: nil, cmpFunc: lambda}
}

// Put into a BST using given comparator
func (bst *BSTMap) Put(key interface{}, val interface{}) {
	if bst.tree == nil {
		// initial insertion
		bst.tree = newNode(key, val)
		bst.size++
	} else {
		// binary search descend
		current := bst.tree
		prev := current
		flag := left
		for current != nil {
			res := bst.cmpFunc(current.key, key)
			switch res {
			case compare.Less: // val > current.val
				prev = current
				current = current.rchild
				flag = right
			case compare.Greater: // val <= current.val
				prev = current
				current = current.lchild
				flag = left
			case compare.Equal:
				// modify value directly.
				current.val = val
				return
			}
		}
		if flag == left {
			prev.lchild = newNode(key, val)
			bst.size++
		} else if flag == right {
			prev.rchild = newNode(key, val)
			bst.size++
		} else {
			panic("unreachable situation.")
		}
	}
}

// Get certain key in BSTTreeMap.
func (bst *BSTMap) Get(key interface{}) interface{} {
	_, c, _ := bst.keySearch(key)
	if c == nil {
		return nil
	}
	return c.val
}

// Delete certain key.
func (bst *BSTMap) Delete(key interface{}) {
	prev, current, isleft := bst.keySearch(key)
	if current == nil {
		// trying to delete an unexist target.
		panic("trying to delete an unexisting key.")
	} else {
		doDelete(prev, current, bst.cmpFunc, isleft)
		bst.size--
		if bst.size == 0 {
			bst.tree.free()
			bst.tree = nil
		}
	}
}

// Size of a BSTMap.
func (bst *BSTMap) Size() int {
	return bst.size
}

// Height calculates the maximum height of a given bst tree.
func (bst *BSTMap) Height() int {
	return height(bst.tree)
}

func height(n *TNode) int {
	if n == nil {
		return 0
	}
	left := height(n.lchild)
	right := height(n.rchild)
	return common.MaxInt(left, right) + 1
}

func doDelete(prev *TNode, cur *TNode, compFunc compare.CompareFunc, isleft bool) {
	if cur == nil {
		return
	}
	if cur.lchild == nil && cur.rchild == nil {
		cur.free()
		if prev != nil {
			if isleft {
				prev.lchild = nil
			} else {
				prev.rchild = nil
			}
		}
		return
	}
	flag := compare.Equal
	if cur.lchild == nil {
		flag = right
	} else if cur.rchild == nil {
		flag = left
	} else {
		cmp := compFunc(cur.lchild.val, cur.rchild.val)
		if cmp == compare.Greater {
			flag = left
		} else if cmp == compare.Less {
			flag = right
		} else {
			panic("unexpected situation.")
		}
	}
	// Do delete
	if flag == left {
		cur.key = cur.lchild.key
		cur.val = cur.lchild.val
		doDelete(cur, cur.lchild, compFunc, true)
	} else {
		cur.key = cur.rchild.key
		cur.val = cur.rchild.val
		doDelete(cur, cur.rchild, compFunc, false)
	}
}

func (bst *BSTMap) keySearch(key interface{}) (prev *TNode, current *TNode, isleft bool) {
	return keySearch(bst.tree, key, bst.cmpFunc)
}

// PreOrderTraverse the given binary tree with a consumer.
func (bst *BSTMap) PreOrderTraverse(fn alias.P1Consumer) {
	bst.doPreOrderTraverse(bst.tree, fn)
}

// InOrderTraverse the given binary tree with a consumer.
func (bst *BSTMap) InOrderTraverse(fn alias.P1Consumer) {
	bst.doInOrderTraverse(bst.tree, fn)
}

// PostOrderTraverse the given binary tree with a consumer.
func (bst *BSTMap) PostOrderTraverse(fn alias.P1Consumer) {
	bst.doPostOrderTraverse(bst.tree, fn)
}

// LevelOrderTraverse the given binary tree with a consumer.
func (bst *BSTMap) LevelOrderTraverse(fn alias.P1Consumer) {
	var queue list.Queue = list.NewSingleLinkedList()
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

func (bst *BSTMap) doPreOrderTraverse(node *TNode, fn alias.P1Consumer) {
	if node == nil {
		return
	}
	fn(node.key)
	bst.doPreOrderTraverse(node.lchild, fn)
	bst.doPreOrderTraverse(node.rchild, fn)
}

func (bst *BSTMap) doInOrderTraverse(node *TNode, fn alias.P1Consumer) {
	if node == nil {
		return
	}
	bst.doInOrderTraverse(node.lchild, fn)
	fn(node.key)
	bst.doInOrderTraverse(node.rchild, fn)
}

func (bst *BSTMap) doPostOrderTraverse(node *TNode, fn alias.P1Consumer) {
	if node == nil {
		return
	}
	bst.doPostOrderTraverse(node.lchild, fn)
	bst.doPostOrderTraverse(node.rchild, fn)
	fn(node.key)
}
