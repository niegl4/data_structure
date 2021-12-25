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
	if e <= s {
		return true
	}

	pivot := s
	root := postOrder[e]
	//下面这段最烧脑
	for i := s; i < e; i++ {
		if i == s {
			if root < postOrder[i] {
				pivot = -1
				break
			}
		}
		if i == e-1 {
			if root > postOrder[i] {
				pivot = e - 1
				break
			}
		}
		if root > postOrder[i] && root < postOrder[i+1] {
			pivot = i
			break
		}
	}

	for i := s; i <= pivot; i++ {
		if postOrder[i] > root {
			return false
		}
	}
	for i := pivot + 1; i < e; i++ {
		if postOrder[i] < root {
			return false
		}
	}
	return verifySequenceOfBstPostOrderCore(postOrder, s, pivot) && verifySequenceOfBstPostOrderCore(postOrder, pivot+1, e-1)
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
	if e <= s {
		return true
	}

	pivot := s
	root := preOrder[s]
	//下面这段最烧脑
	for i := s + 1; i <= e; i++ {
		if i == s+1 {
			if root < preOrder[i] {
				pivot = 0 //【注意：这里从后序的-1，改为了0】
				break
			}
		}
		if i == e {
			if root > preOrder[i] {
				pivot = e
				break
			}
		}
		if root > preOrder[i] && root < preOrder[i+1] {
			pivot = i
			break
		}
	}

	for i := s + 1; i <= pivot; i++ {
		if preOrder[i] > root {
			return false
		}
	}
	for i := pivot + 1; i < e; i++ {
		if preOrder[i] < root {
			return false
		}
	}
	return verifySequenceOfBstPreOrderCore(preOrder, s+1, pivot) && verifySequenceOfBstPreOrderCore(preOrder, pivot+1, e)
}
