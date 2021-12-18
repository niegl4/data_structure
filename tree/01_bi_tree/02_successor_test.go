package _1_bi_tree

import (
	"data_structure/tree/common"
	"testing"
)


func TestGetSuccessor(t *testing.T) {
	root, err := constructTmp([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
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


	node := getSuccessor(root) //有右子树的情况 [5]
	t.Log(node)
	node = getSuccessor(root.lChild) //没有右子树，但它是父节点的左子节点 [1]
	t.Log(node)
	node = getSuccessor(root.lChild.lChild.rChild) //没有右子树，它是父节点的右子节点 [2]
	t.Log(node)
	node = getSuccessor(root.rChild.rChild) //没有右子树，它是父节点的右子节点 nil
	t.Log(node)
}

