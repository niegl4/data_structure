package _1_bi_tree

/*
六十八-1
二叉排序树中，寻找两节点的最近公共父节点。
todo
*/

/*
六十八-2
树中节点有指向父节点的指针，寻找两节点的最近公共父节点。

见单链表的09_first_common_node.go
*/

/*
六十八-3
树中节点没有指向父节点的指针，寻找两节点的最近公共父节点。
****
*/
func getNodePath(root *BiTree, targetNode *BiTree, path *[]*BiTree) bool {
	if root == targetNode {
		return true
	}
	*path = append(*path, root)
	found := false

	iterator := make([]*BiTree, 0)
	if root.lChild != nil {
		iterator = append(iterator, root.lChild)
	}
	if root.rChild != nil {
		iterator = append(iterator, root.rChild)
	}
	for i := 0; i < len(iterator) && !found; i++ {
		found = getNodePath(iterator[i], targetNode, path)
	}

	if !found {
		*path = (*path)[:len(*path)-1]
	}
	return found
}

func getLastCommonNode(path1, path2 []*BiTree) *BiTree {
	length1 := len(path1)
	length2 := len(path2)
	length := length1
	if length2 < length1 {
		length = length2
	}
	var node *BiTree
	for i := 0; i < length; i++ {
		if path1[i] == path2[i] {
			node = path1[i]
		}
	}
	return node
}

func getLastCommonParent(root, node1, node2 *BiTree) *BiTree {
	if root == nil || node1 == nil || node2 == nil {
		return nil
	}

	path1 := make([]*BiTree, 0)
	getNodePath(root, node1, &path1)

	path2 := make([]*BiTree, 0)
	getNodePath(root, node2, &path2)

	return getLastCommonNode(path1, path2)
}
