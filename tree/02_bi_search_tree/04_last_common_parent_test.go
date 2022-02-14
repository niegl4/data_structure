package _2_bi_search_tree

import "testing"

/*
       5
     3   7
   2  4 6  8
*/
func TestLastCommonParent(t *testing.T) {
	root := &TreeNode{
		data: 5,
		lChild: &TreeNode{
			data: 3,
			lChild: &TreeNode{
				data:   2,
				lChild: nil,
				rChild: nil,
				parent: nil,
			},
			rChild: &TreeNode{
				data:   4,
				lChild: nil,
				rChild: nil,
				parent: nil,
			},
			parent: nil,
		},
		rChild: &TreeNode{
			data: 7,
			lChild: &TreeNode{
				data:   6,
				lChild: nil,
				rChild: nil,
				parent: nil,
			},
			rChild: &TreeNode{
				data:   8,
				lChild: nil,
				rChild: nil,
				parent: nil,
			},
			parent: nil,
		},
		parent: nil,
	}

	node := lastCommonParent(root, root.rChild.lChild, root.rChild.rChild)
	if node == nil {
		t.Log(node)
	} else {
		t.Log(node.data)
	}

	node = lastCommonParent(root, root.lChild.lChild, root.rChild.rChild)
	if node == nil {
		t.Log(node)
	} else {
		t.Log(node.data)
	}

}