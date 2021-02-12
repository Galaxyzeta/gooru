package testds

import (
	"strconv"
	"testing"

	"galaxyzeta.com/algo/compare"
	"galaxyzeta.com/ds/tree"
	"galaxyzeta.com/util/assert"
)

func TestBST(t *testing.T) {
	// == put test
	bst := tree.NewBST(compare.BasicCompare)
	bst.Put(1, 10086)
	bst.Put(2, 2)
	bst.Put(3, 3)
	bst.Put(4, 4)
	bst.Put(5, 5)
	bst.Put(1, 1)
	assert.EQ(bst.Size(), 5)
	assert.EQ(bst.Height(), 5)
	// == traverse test
	want := "123451234554321"
	actual := ""
	bst.PreOrderTraverse(func(param interface{}) {
		actual += strconv.Itoa(param.(int))
	})
	bst.InOrderTraverse(func(param interface{}) {
		actual += strconv.Itoa(param.(int))
	})
	bst.PostOrderTraverse(func(param interface{}) {
		actual += strconv.Itoa(param.(int))
	})
	assert.EQ(actual, want)
	// == deletion test
	assert.EQ(bst.Delete(1), 1)
	assert.EQ(bst.Get(1), nil)
	assert.EQ(bst.Get(2), 2)
	assert.EQ(bst.Delete(3), 3)
	assert.EQ(bst.Get(2), 2)
	assert.EQ(bst.Get(3), nil)
	assert.EQ(bst.Get(4), 4)
	assert.EQ(bst.Get(5), 5)
	assert.EQ(bst.Delete(4), 4)
	assert.EQ(bst.Delete(5), 5)
	assert.EQ(bst.Height(), 1)
	assert.EQ(bst.Delete(2), 2)
	assert.EQ(bst.Get(2), nil)
	assert.EQ(bst.Height(), 0)
}
