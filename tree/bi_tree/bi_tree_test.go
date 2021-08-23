package bi_tree

import (
	"data_structure/tree/common"
	"testing"
)

func TestBiTree_PreOrderTraverse(t *testing.T) {
	var biTestOp common.Operate = func(data interface{}) () {
		t.Log(data)
	}
	t.Log("BiTree 前序遍历")
	BiTreeRoot.PreOrderTraverse(biTestOp)
	t.Log("\n")

	t.Log("BiTree 中序遍历")
	BiTreeRoot.InOrderTraverse(biTestOp)
	t.Log("\n")

	t.Log("BiTree 后序遍历")
	BiTreeRoot.PostOrderTraverse(biTestOp)
}
