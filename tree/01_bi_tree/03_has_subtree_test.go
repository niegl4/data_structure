package _1_bi_tree

import "testing"

func TestTree1HasTree2(t *testing.T) {
	/*
					8
				  /  \
				8	  7
		       / \
			  9	  2
		     / \
		    4	7
	*/
	t1Node1 := &BiTree{Data: 8}
	t1Node2 := &BiTree{Data: 8}
	t1Node3 := &BiTree{Data: 7}
	t1Node4 := &BiTree{Data: 9}
	t1Node5 := &BiTree{Data: 2}
	t1Node6 := &BiTree{Data: 4}
	t1Node7 := &BiTree{Data: 7}
	t1Node1.lChild = t1Node2
	t1Node1.rChild = t1Node3
	t1Node2.lChild = t1Node4
	t1Node2.rChild = t1Node5
	t1Node5.lChild = t1Node6
	t1Node5.rChild = t1Node7

	/*
				8
		       / \
			  9	  2
	*/
	t2Node1 := &BiTree{Data: 8}
	t2Node2 := &BiTree{Data: 9}
	t2Node3 := &BiTree{Data: 2}
	t2Node1.lChild = t2Node2
	t2Node1.rChild = t2Node3

	t.Log(tree1HasTree2(t1Node1, t2Node1))
}
