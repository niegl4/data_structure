package stack

import "data_structure/common"

/*
三十一
*
输入一个入栈序列，一个出栈序列，判断该出栈序列是否是该栈的一个弹出序列。
已知入栈序列的所有数字都不等。

如果下一个弹出的数字刚好是辅助栈的栈顶数字，直接弹出，接着判断下一个弹出数字。
如果不是辅助栈的栈顶数字，就把压栈序列中还没入栈的数字压入辅助栈，直到把下一个需要弹出的数字压入辅助栈为止。
	如果在压栈序列终止前，找到了目标数字，则继续循环这个过程。
	如果在压栈序列终止前，没找到目标数字，返回false。
*/
func isPopOrder(push, pop []int) bool {
	if len(push) == 0 || len(pop) == 0 {
		return false
	}

	pushIdx := 0
	pushIdxMax := len(push) - 1
	supStack := common.NewStack()
	for _, popNum := range pop {
		if supStack.Len() == 0 {
			res, pushIdxNew := isPopNum(push, pushIdx, pushIdxMax, popNum, supStack)
			if !res {
				return false
			} else {
				pushIdx = pushIdxNew
			}
		} else {
			stackTopNum, _ := supStack.Pop()
			if stackTopNum == popNum {
				continue
			} else {
				supStack.Push(stackTopNum)
				res, pushIdxNew := isPopNum(push, pushIdx, pushIdxMax, popNum, supStack)
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

func isPopNum(push []int, pushIdx, pushIdxMax, popNum int, supStack *common.Stack) (bool, int) {
	if pushIdx > pushIdxMax {
		return false, 0
	}
	for ; pushIdx <= pushIdxMax; pushIdx++ {
		supStack.Push(push[pushIdx])
		if push[pushIdx] == popNum {
			break
		}
	}
	stackTopNum, _ := supStack.Pop()
	if stackTopNum != popNum {
		return false, 0
	} else {
		pushIdx++
		return true, pushIdx
	}
}
