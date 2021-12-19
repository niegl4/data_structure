package common

import (
	"container/list"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.In(1)
	q.In(2)
	q.In(3)
	t.Log(q.Out())
	t.Log(q.Out())
	t.Log(q.Out())
	t.Log(q.Out())
}

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
