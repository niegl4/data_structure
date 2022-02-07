package stack

import (
	"data_structure/common"
	"errors"
)

/*
三十
定义栈的数据结构。使得它有一个min方法，可以找到栈中的最小值。min，push，pop方法的时间复杂度都是O(1)。
*/

type stackWithMin struct {
	dataStack *common.Stack
	minStack  *common.Stack
}

func NewStackWithMin() *stackWithMin {
	return &stackWithMin{
		dataStack: common.NewStack(),
		minStack:  common.NewStack(),
	}
}

func (minS *stackWithMin) min() (int, error) {
	if minS == nil {
		return 0, errors.New("stack is nil")
	}
	if minS.minStack.Len() == 0 {
		return 0, errors.New("stack is empty")
	}
	val, err := minS.minStack.Pop()
	if err != nil {
		return 0, err
	}
	minS.minStack.Push(val)
	return val, nil
}

func (minS *stackWithMin) push(val int) error {
	if minS == nil {
		return errors.New("stack is nil")
	}
	if minS.minStack.Len() == 0 {
		minS.minStack.Push(val)
	} else {
		curMin, _ := minS.minStack.Pop()
		minS.minStack.Push(curMin)
		if val < curMin {
			minS.minStack.Push(val)
		} else {
			minS.minStack.Push(curMin)
		}
	}
	minS.dataStack.Push(val)
	return nil
}

func (minS *stackWithMin) pop() (int, error) {
	if minS == nil {
		return 0, errors.New("stack is nil")
	}
	if minS.minStack.Len() == 0 {
		return 0, errors.New("stack is empty")
	}
	val, _ := minS.dataStack.Pop()
	_, _ = minS.minStack.Pop()
	return val, nil
}
