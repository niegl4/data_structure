package common

import (
	"container/list"
	"errors"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{
		list: list.New(),
	}
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Push(val int) {
	s.list.PushBack(val)
}

func (s *Stack) Pop() (val int, err error) {
	if s.list.Len() == 0 {
		return 0, errors.New("stack is empty")
	}
	ele := s.list.Back()
	s.list.Remove(ele)
	val = ele.Value.(int)
	return val, nil
}
