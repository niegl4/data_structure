package stack

import "data_structure/common"

/*
三十一
输入一个入栈序列，一个出栈序列，判断该出栈序列是否是该栈的一个弹出序列。
已知入栈序列的所有数字都不等。
 */
func isPopOrder(push, pop []int) bool {
	if len(push) == 0 || len(pop) == 0 {
		return false
	}

	pushIdx := 0
	pushIdxMax := len(push) - 1
	stack := common.NewStack()
	for _, popNum := range pop {
		if stack.Len() == 0 {
			res, pushIdxNew := isPopNum(push, pushIdx, pushIdxMax, popNum, stack)
			if !res {
				return false
			} else {
				pushIdx = pushIdxNew
			}
		} else {
			stackTopNum, _ := stack.Pop()
			if stackTopNum == popNum {
				continue
			} else {
				stack.Push(stackTopNum)
				res, pushIdxNew := isPopNum(push, pushIdx, pushIdxMax, popNum, stack)
				if !res {
					return false
				} else {
					pushIdx = pushIdxNew
				}
			}
		}
	}
	return true
}

func isPopNum(push []int, pushIdx, pushIdxMax, popNum int, stack *common.Stack) (bool, int) {
	if pushIdx > pushIdxMax {
		return false, 0
	}
	for ; pushIdx <= pushIdxMax; pushIdx++ {
		stack.Push(push[pushIdx])
		if push[pushIdx] == popNum {
			break
		}
	}
	stackTopNum, _ := stack.Pop()
	if stackTopNum != popNum {
		return false, 0
	} else {
		pushIdx++
		return true, pushIdx
	}
}
