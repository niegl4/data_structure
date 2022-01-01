package _1_bi_tree

/*
五十五-1
二叉树的深度。

五十五-2，见01_judge_balance.go
*/
func depth(root *BiTree) int {
	if root == nil {
		return 0
	}
	return depthCore(root)
}

func depthCore(node *BiTree) int {
	if node == nil {
		return 0
	}
	left := depthCore(node.lChild)
	right := depthCore(node.rChild)
	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}
