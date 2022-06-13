package _1_bi_tree

import "testing"

func TestMirrorBiTree(t *testing.T) {
	/*
		         1
			   /  \
			  2    3
		     / \  / \
			4  5 6  7
	*/
	t1Node1 := &BiTree{Data: 1}
	t1Node2 := &BiTree{Data: 2}
	t1Node3 := &BiTree{Data: 3}
	t1Node4 := &BiTree{Data: 4}
	t1Node5 := &BiTree{Data: 5}
	t1Node6 := &BiTree{Data: 6}
	t1Node7 := &BiTree{Data: 7}
	t1Node1.lChild = t1Node2
	t1Node1.rChild = t1Node3
	t1Node2.lChild = t1Node4
	t1Node2.rChild = t1Node5
	t1Node3.lChild = t1Node6
	t1Node3.rChild = t1Node7
	t.Log(mirrorBiTreeV2(t1Node1).levelOrderTraverse1())

	mirrorBiTree(t1Node1)
	t.Log(t1Node1.levelOrderTraverse1())
}

func TestIsSymmetrical(t *testing.T) {
	/*
		         1
			   /  \
			  2    2
		     / \  / \
			4  5 5  4
	*/
	t1Node1 := &BiTree{Data: 1}
	t1Node2 := &BiTree{Data: 2}
	t1Node3 := &BiTree{Data: 2}
	t1Node4 := &BiTree{Data: 4}
	t1Node5 := &BiTree{Data: 5}
	t1Node6 := &BiTree{Data: 5}
	t1Node7 := &BiTree{Data: 4}
	t1Node1.lChild = t1Node2
	t1Node1.rChild = t1Node3
	t1Node2.lChild = t1Node4
	t1Node2.rChild = t1Node5
	t1Node3.lChild = t1Node6
	t1Node3.rChild = t1Node7
	t.Log(isSymmetrical(t1Node1))
	/*
		         1
			   /  \
			  1    1
		     / \  /
			1  1 1
	*/
	t2Node1 := &BiTree{Data: 1}
	t2Node2 := &BiTree{Data: 1}
	t2Node3 := &BiTree{Data: 1}
	t2Node4 := &BiTree{Data: 1}
	t2Node5 := &BiTree{Data: 1}
	t2Node6 := &BiTree{Data: 1}
	t2Node1.lChild = t2Node2
	t2Node1.rChild = t2Node3
	t2Node2.lChild = t2Node4
	t2Node2.rChild = t2Node5
	t2Node3.lChild = t2Node6
	t.Log(isSymmetrical(t2Node1))
}
