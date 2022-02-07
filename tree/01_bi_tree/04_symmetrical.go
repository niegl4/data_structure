package _1_bi_tree

/*
二十八
判断一颗二叉树是否是对称的。
*/
func isSymmetrical(root *BiTree) bool {
	if root == nil {
		return false
	}
	return isSymmetricalCore(root, root)
}

func isSymmetricalCore(node1, node2 *BiTree) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil && node2 != nil {
		return false
	} else if node1 != nil && node2 == nil {
		return false
	}
	if node1 != nil && node2 != nil && (node1.Data).(int) != (node2.Data).(int) {
		return false
	}
	return isSymmetricalCore(node1.lChild, node2.rChild) && isSymmetricalCore(node1.rChild, node2.lChild)
}
