package _1_bi_tree

import "testing"

func TestBiTree_LevelOrderTraverse(t *testing.T) {
	t.Log("BiTree 层序遍历")
	tmp := BiTreeRoot.levelOrderTraverse1()
	t.Log(tmp)

	t.Log("BiTree 层序遍历")
	levels := BiTreeRoot.levelOrderTraverse2()
	for _, level := range levels {
		t.Log(level)
	}

	t.Log("BiTree 层序遍历")
	levels = BiTreeRoot.levelOrderTraverse3()
	for _, level := range levels {
		t.Log(level)
	}
}
