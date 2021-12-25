package _1_bi_tree

import (
	"testing"
)

func TestBiTree_PreOrderTraverse(t *testing.T) {
	t.Log("BiTree 前序遍历")
	tmp := BiTreeRoot.PreOrderTraverse()
	t.Logf("%v", tmp)

	t.Log("BiTree 中序遍历")
	tmp = BiTreeRoot.InOrderTraverse()
	t.Logf("%v", tmp)

	t.Log("BiTree 后序遍历")
	tmp = BiTreeRoot.PostOrderTraverse()
	t.Logf("%v", tmp)

	t.Log("BiTree 层序遍历")
	levelRes := BiTreeRoot.LevelOrderTraverse()
	t.Logf("%v", levelRes)
}
