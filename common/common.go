package common

import (
	"container/list"
	"errors"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{
		list: list.New(),
	}
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func (q *Queue) In(val int) {
	q.list.PushBack(val)
}

func (q *Queue) Out() (val int, err error) {
	if q.list.Len() == 0 {
		return 0, errors.New("queue is empty")
	}
	ele := q.list.Front()
	q.list.Remove(ele)
	val = ele.Value.(int)
	return val, nil
}

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


