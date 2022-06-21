package _2_bi_search_tree

/*
三十三-1
判断整数数组是不是某个二叉搜索树的后续遍历结果。数组中的每个数字都不相等。
*/
func verifySequenceOfBstPostOrder(postOrder []int) bool {
	length := len(postOrder)
	if length <= 1 {
		return true
	}

	//校验数组中的元素，是否有相等的
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if postOrder[i] == postOrder[j] {
				return false
			}
		}
	}
	return verifySequenceOfBstPostOrderCore(postOrder, 0, length-1)
}

func verifySequenceOfBstPostOrderCore(postOrder []int, s, e int) bool {
	if s >= e {
		return true
	}

	pivot := s
	root := postOrder[e]
	for ; pivot < e; pivot++ {
		if postOrder[pivot] > root {
			break
		}
	}

	j := pivot
	for ; j < e; j++ {
		if postOrder[j] < root {
			return false
		}
	}
	return verifySequenceOfBstPostOrderCore(postOrder, s, pivot-1) && verifySequenceOfBstPostOrderCore(postOrder, pivot, e-1)
}

/*
三十三-2
判断整数数组是不是某个二叉搜索树的前续遍历结果。数组中的每个数字都不相等。
*/
func verifySequenceOfBstPreOrder(preOrder []int) bool {
	length := len(preOrder)
	if length <= 1 {
		return true
	}

	//校验数组中的元素，是否有相等的
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if preOrder[i] == preOrder[j] {
				return false
			}
		}
	}
	return verifySequenceOfBstPreOrderCore(preOrder, 0, length-1)
}

func verifySequenceOfBstPreOrderCore(preOrder []int, s, e int) bool {
	if s >= e {
		return true
	}

	root := preOrder[s]
	pivot := s + 1
	for ; pivot <= e; pivot++ {
		if preOrder[pivot] > root {
			break
		}
	}

	j := pivot
	for ; j <= e; j++ {
		if preOrder[j] < root {
			return false
		}
	}
	return verifySequenceOfBstPreOrderCore(preOrder, s+1, pivot-1) && verifySequenceOfBstPreOrderCore(preOrder, pivot, e)
}
