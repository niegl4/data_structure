package _1_bi_tree

/*
三十四
*
输入一棵二叉树和一个整数，返回二叉树中节点值的和为该整数的所有路径。

不仅需要一个容器保存所有path，还需要一个辅助栈保存当前路径。
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

//递归核心函数，通过int类型传递curSum，返回递归上一级时，不用处理curSum；通过*[]类型传递path和res，返回上一级时，需要截掉当前节点
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
