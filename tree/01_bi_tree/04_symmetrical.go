package _1_bi_tree

/*
二十七
输入一棵二叉树，输出它的镜像。

两棵树互为镜像，与"二十八 一棵树是否对称"不同。
*/
func mirrorBiTree(root *BiTree) {
	if root == nil {
		return
	}
	if root.lChild == nil && root.rChild == nil {
		return
	}

	tmp := root.lChild
	root.lChild = root.rChild
	root.rChild = tmp

	mirrorBiTree(root.lChild)
	mirrorBiTree(root.rChild)
}

/*
二十八
判断一颗二叉树是否是对称的。
*
递归核心函数，就是两个节点的值进行判等操作。
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
