package stack

import (
	"data_structure/queue"
	"errors"
)

/*
九-2
两个队列实现一个栈
*/

type stackTmp struct {
	queue1, queue2 *queue.Queue
}

func (s *stackTmp) push(val int) error {
	if s == nil {
		return errors.New("need init stack pointer")
	}
	if s.queue1 == nil {
		s.queue1 = queue.NewQueue()
	}
	s.queue1.In(val)
	return nil
}

func (s *stackTmp) pop() (int, error) {
	if s == nil {
		return 0, errors.New("need init stack pointer")
	}
	if s.queue1 == nil {
		s.queue1 = queue.NewQueue()
	}
	if s.queue2 == nil {
		s.queue2 = queue.NewQueue()
	}

	if s.queue1.Len() == 0 && s.queue2.Len() == 0 {
		return 0, errors.New("stack is empty")
	}
	if s.queue1.Len() > 0 {
		for s.queue1.Len() >= 2 {
			val, err := s.queue1.Out()
			if err != nil {
				return 0, err
			}
			s.queue2.In(val)
		}
		return s.queue1.Out()
	} else {
		for s.queue2.Len() >= 2 {
			val, err := s.queue2.Out()
			if err != nil {
				return 0, err
			}
			s.queue1.In(val)
		}
		return s.queue2.Out()
	}
}
