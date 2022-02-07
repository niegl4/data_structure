package stack

import (
	"data_structure/common"
	"errors"
)

/*
九-2
两个队列实现一个栈
*/

type stackTmp struct {
	queue1, queue2 *common.Queue
}

func (s *stackTmp) push(val int) error {
	if s == nil {
		return errors.New("need init stack pointer")
	}
	if s.queue1 == nil {
		s.queue1 = common.NewQueue()
	}
	s.queue1.Enqueue(val)
	return nil
}

func (s *stackTmp) pop() (int, error) {
	if s == nil {
		return 0, errors.New("need init stack pointer")
	}
	if s.queue1 == nil {
		s.queue1 = common.NewQueue()
	}
	if s.queue2 == nil {
		s.queue2 = common.NewQueue()
	}

	if s.queue1.Len() == 0 && s.queue2.Len() == 0 {
		return 0, errors.New("stack is empty")
	}
	if s.queue1.Len() > 0 {
		for s.queue1.Len() >= 2 {
			val, err := s.queue1.Dequeue()
			if err != nil {
				return 0, err
			}
			s.queue2.Enqueue(val)
		}
		return s.queue1.Dequeue()
	} else {
		for s.queue2.Len() >= 2 {
			val, err := s.queue2.Dequeue()
			if err != nil {
				return 0, err
			}
			s.queue1.Enqueue(val)
		}
		return s.queue2.Dequeue()
	}
}
