package queue

import (
	"container/list"
	"errors"
)

type queue struct {
	list *list.List
}

func NewQueue() *queue {
	return &queue{
		list: list.New(),
	}
}

func (q *queue) Len() int {
	return q.list.Len()
}

func (q *queue) In(val int) {
	q.list.PushBack(val)
}

func (q *queue) Out() (val int, err error) {
	if q.list.Len() == 0 {
		return 0, errors.New("queue is empty")
	}
	ele := q.list.Front()
	q.list.Remove(ele)
	val = ele.Value.(int)
	return val, nil
}
