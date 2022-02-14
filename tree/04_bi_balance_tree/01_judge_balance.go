package _4_bi_balance_tree

type BiTree struct {
	Data   interface{}
	LChild *BiTree
	RChild *BiTree
}

/*
五十五-1，见08_depth.go

五十五-2
判断二叉树的平衡性。任意节点的左，右子树的深度差不超过1，即为平衡二叉树。
*/
func judgeBalance(root *BiTree) bool {
	if root == nil {
		return true
	}
	depth := 0
	return judgeBalanceCore(root, &depth)
}

func judgeBalanceCore(node *BiTree, depth *int) bool {
	if node == nil {
		*depth = 0
		return true
	}

	lDep, rDep := 0, 0
	lBal := judgeBalanceCore(node.LChild, &lDep)
	rBal := judgeBalanceCore(node.RChild, &rDep)
	if lBal && rBal { //左右子树都平衡
		diff := lDep - rDep
		if diff >= -1 && diff <= 1 { //左右子树的高度小于等于1
			if lDep >= rDep {
				*depth = lDep + 1
			} else {
				*depth = rDep + 1
			}
			return true
		}
	}
	return false
}
