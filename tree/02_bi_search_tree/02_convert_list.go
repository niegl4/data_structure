package _2_bi_search_tree

/*
三十六
***
输入一棵二叉搜索树，转换成一个排序的双向链表。不能创建新节点，只能调整树中节点的指向。
*/
func bstConvert(root *TreeNode) (head *TreeNode) {
	if root == nil {
		return nil
	}
	var lastNodeInList *TreeNode
	bstConvertCore(root, &lastNodeInList)

	for lastNodeInList != nil && lastNodeInList.lChild != nil {
		lastNodeInList = lastNodeInList.lChild
	}
	head = lastNodeInList
	return head
}

func bstConvertCore(node *TreeNode, lastNodeInList **TreeNode) {
	if node == nil {
		return
	}

	if node.lChild != nil {
		bstConvertCore(node.lChild, lastNodeInList)
	}

	node.lChild = *lastNodeInList
	if *lastNodeInList != nil {
		(*lastNodeInList).rChild = node
	}
	*lastNodeInList = node

	if node.rChild != nil {
		bstConvertCore(node.rChild, lastNodeInList)
	}
}
