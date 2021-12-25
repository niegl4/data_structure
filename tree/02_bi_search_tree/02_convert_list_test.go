package _2_bi_search_tree

import "testing"

func TestBstConvert(t *testing.T) {
	root := &TreeNode{
		data: 10,
		lChild: &TreeNode{
			data: 6,
			lChild: &TreeNode{
				data:   4,
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
		rChild: &TreeNode{
			data: 14,
			lChild: &TreeNode{
				data:   12,
				lChild: nil,
				rChild: nil,
				parent: nil,
			},
			rChild: &TreeNode{
				data:   16,
				lChild: nil,
				rChild: nil,
				parent: nil,
			},
			parent: nil,
		},
		parent: nil,
	}
	head := bstConvert(root)
	node := head
	tmp := make([]float64, 0)
	for node != nil {
		tmp = append(tmp, node.data)
		node = node.rChild
	}
	t.Log(tmp)
}
