package __r_b_tree

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
