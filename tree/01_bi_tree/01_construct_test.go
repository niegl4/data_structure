package _1_bi_tree

import (
	"data_structure/tree/common"
	"testing"
)

func TestConstruct(t *testing.T) {
	root, err := construct([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	if err != nil {
		t.Fatal(err)
	}
	tmp := make([]interface{}, 0, 16)
	var biTestOp common.Operate = func(data interface{}) () {
		tmp = append(tmp, data)
	}
	root.PreOrderTraverse(biTestOp)
	t.Log(tmp)

	tmp = tmp[0:0]
	root.InOrderTraverse(biTestOp)
	t.Log(tmp)

	tmp = tmp[0:0]
	root.PostOrderTraverse(biTestOp)
	t.Log(tmp)
}
