package stack

import (
	"container/list"
	"errors"
)

type stack struct {
	list *list.List
}

func NewStack() *stack {
	return &stack{
		list: list.New(),
	}
}

func (s *stack) Len() int {
	return s.list.Len()
}

func (s *stack) Push(val int) {
	s.list.PushBack(val)
}

func (s *stack) Pop() (val int, err error) {
	if s.list.Len() == 0 {
		return 0, errors.New("stack is empty")
	}
	ele := s.list.Back()
	s.list.Remove(ele)
	val = ele.Value.(int)
	return val, nil
}
