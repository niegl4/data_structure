package _1_bi_tree

/*
三十四
*
输入一棵二叉树和一个整数，返回二叉树中节点值的和为该整数的所有路径。
*/
func findPath(root *BiTree, expectNum int) [][]*BiTree {
	if root == nil {
		return nil
	}
	curSum := 0
	path := make([]*BiTree, 0) //实际上是一个栈
	res := make([][]*BiTree, 0)
	findPathCore(root, expectNum, curSum, &path, &res)
	return res
}

func findPathCore(node *BiTree, expectNum, curSum int, path *[]*BiTree, res *[][]*BiTree) {
	curSum += (node.Data).(int)
	*path = append(*path, node)

	isLeaf := node.lChild == nil && node.rChild == nil
	if curSum == expectNum && isLeaf {
		var tmp []*BiTree
		tmp = append(tmp, *path...)
		*res = append(*res, tmp)
	}
	if curSum < expectNum {
		if node.lChild != nil {
			findPathCore(node.lChild, expectNum, curSum, path, res)
		}
		if node.rChild != nil {
			findPathCore(node.rChild, expectNum, curSum, path, res)
		}
	}

	*path = (*path)[0 : len(*path)-1]
}
