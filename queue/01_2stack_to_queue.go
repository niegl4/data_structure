package queue

import (
	"data_structure/stack"
	"errors"
)

/*
九-1
两个栈实现一个队列
*/

type queueTmp struct {
	stack1, stack2 *stack.Stack
}

func (q *queueTmp) in(val int) error {
	if q == nil {
		return errors.New("need init queue pointer")
	}
	if q.stack1 == nil {
		q.stack1 = stack.NewStack()
	}

	q.stack1.Push(val)
	return nil
}

func (q *queueTmp) out() (int, error) {
	if q == nil {
		return 0, errors.New("need init queue pointer")
	}
	if q.stack1 == nil {
		q.stack1 = stack.NewStack()
	}
	if q.stack2 == nil {
		q.stack2 = stack.NewStack()
	}

	if q.stack1.Len()+q.stack2.Len() == 0 {
		return 0, errors.New("queue is empty")
	}
	if q.stack2.Len() > 0 { //stack2不为空，直接弹出
		return q.stack2.Pop()
	} else { //stack2为空，stack1逐个弹出，压入stack2
		for q.stack1.Len() > 0 {
			val, err := q.stack1.Pop()
			if err != nil {
				return 0, err
			}
			q.stack2.Push(val)
		}
		return q.stack2.Pop()
	}
}
