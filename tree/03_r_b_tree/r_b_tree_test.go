package _3_r_b_tree

import (
	"testing"
)

func TestRBTree_GetDeepth(t *testing.T) {
	var rbt RBTree
	rbt.Add(9)
	rbt.Add(8)
	rbt.Add(7)
	rbt.Add(6)
	rbt.Add(5)
	rbt.Add(4)
	rbt.Add(3)
	rbt.Add(2)
	rbt.Add(1)
	t.Log(rbt.GetDeepth())
}

func TestRBTree_Delete(t *testing.T) {
	var (
		rbtOp = func(data interface{}) {
			t.Log(data)
		}
	)
	var rbt RBTree
	rbt.Add(2)
	rbt.Add(1)
	rbt.InOrderTravel(rbtOp)

	//这个例子可以说明，原本的Delete方法有bug
	rbt.Delete(float64(2))
	rbt.InOrderTravel(rbtOp)
}

func TestRBTree_DeleteByMe(t *testing.T) {
	var (
		rbtOp = func(data interface{}) {
			t.Log(data)
		}
	)
	var rbt RBTree
	rbt.Add(2)
	rbt.Add(1)
	rbt.InOrderTravel(rbtOp)

	//这个例子可以说明，原本的Delete方法有bug
	rbt.DeleteByMe(float64(2))
	rbt.InOrderTravel(rbtOp)
}
