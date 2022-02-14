package _2_bi_search_tree

/*
六十八-1
二叉排序树中，寻找两节点的最近公共父节点。

问题转化为：在二叉排序树中查找节点(值在两个节点值之间的父节点，只会有一个)
*/

func lastCommonParent(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if p.data < root.data && q.data < root.data {
		return lastCommonParent(root.lChild, p, q)
	} else if p.data > root.data && q.data > root.data {
		return lastCommonParent(root.rChild, p, q)
	} else {
		return root
	}
}


