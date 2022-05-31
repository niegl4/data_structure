package _2_bi_search_tree

/*
五十四
二叉搜索树的第k大节点。
*

递归核心函数，不太容易想到。
二叉搜索树，中序遍历即为从小到大，第k大节点，需要用中序遍历。（题目不太严谨，应该是从小到大顺序，第k大节点）
通过node是否为nil，来标志是否已经找到目标节点。
通过传递*int，来进行计数。
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
