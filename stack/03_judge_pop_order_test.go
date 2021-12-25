package stack

import "testing"

func TestIsPopOrder(t *testing.T) {
	t.Log(isPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2}))
	t.Log(isPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
