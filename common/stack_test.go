package common

import (
	"container/list"
	"testing"
)

func TestStack(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	stack := NewStack()
	for _, num := range arr {
		stack.Push(num)
	}

	for stack.Len() > 0 {
		t.Log(stack.Pop())
	}

	var l list.List
	l.PushBack(1)
	t.Log(l.Back().Value.(int))
}
