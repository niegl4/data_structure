package _2_bi_search_tree

/*
五十四
二叉搜索树的第k大节点。
*/
func kthNode(root *TreeNode, k int) *TreeNode {
	if root == nil || k < 1 {
		return nil
	}
	return kthNodeCore(root, &k)
}

func kthNodeCore(root *TreeNode, k *int) *TreeNode {
	var node *TreeNode
	if root.lChild != nil {
		node = kthNodeCore(root.lChild, k)
	}

	if node == nil {
		if *k == 1 {
			node = root
		}
		*k--
	}

	if node == nil && root.rChild != nil {
		node = kthNodeCore(root.rChild, k)
	}
	return node
}
