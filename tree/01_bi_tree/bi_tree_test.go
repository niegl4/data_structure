package _1_bi_tree

import (
	"data_structure/tree/common"
	"testing"
)

func TestBiTree_PreOrderTraverse(t *testing.T) {
	tmp := make([]interface{}, 0, 16)
	var biTestOp common.Operate = func(data interface{}) () {
		tmp = append(tmp, data)
	}
	t.Log("BiTree 前序遍历")
	BiTreeRoot.PreOrderTraverse(biTestOp)
	t.Logf("%v", tmp)
	tmp = tmp[0:0]
	t.Log("\n")

	t.Log("BiTree 中序遍历")
	BiTreeRoot.InOrderTraverse(biTestOp)
	t.Logf("%v", tmp)
	tmp = tmp[0:0]
	t.Log("\n")

	t.Log("BiTree 后序遍历")
	BiTreeRoot.PostOrderTraverse(biTestOp)
	t.Logf("%v", tmp)
}
