package _1_bi_tree

/*
二十六
判断树2是不是树1的子结构
*/
func tree1HasTree2(root1, root2 *BiTree) bool {
	if root1 == nil || root2 == nil {
		return false
	}

	result := false
	if (root1.Data).(int) == (root2.Data).(int) {
		result = tree1HasTree2Core(root1, root2)
	}
	if !result {
		result = tree1HasTree2(root1.lChild, root2)
	}
	if !result {
		result = tree1HasTree2(root1.rChild, root2)
	}
	return result
}

func tree1HasTree2Core(node1, node2 *BiTree) bool {
	if node2 == nil {
		return true
	} else if node1 == nil {
		return false
	}
	if node1.Data.(int) != node2.Data.(int) {
		return false
	}
	return tree1HasTree2Core(node1.lChild, node2.lChild) && tree1HasTree2Core(node1.rChild, node2.rChild)
}
