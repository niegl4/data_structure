package _1_bi_tree

/*
二十六
判断树2是不是树1的子结构
*
result控制剩余递归调用的设计，很巧妙
它就是一个前序遍历的变形。
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

//递归核心函数，就是用来判断以node1为根节点的树，是否包含有以node2为根节点的树。
//注意：node1可以是原tree1的任意节点，而不一定是根节点。node2也是这样。
//这样设计递归核心函数，可以最大限度的复用它，而且保证核心函数简单易懂。
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
