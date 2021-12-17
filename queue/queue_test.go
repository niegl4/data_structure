package queue

import "testing"

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
