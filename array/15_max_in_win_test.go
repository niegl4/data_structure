package array

import "testing"

func TestMaxInWindows(t *testing.T) {
	t.Log(maxInWindows([]int{2, 3, 4, 2, 6, 2, 5, 1, 0, -1}, 3))
}
