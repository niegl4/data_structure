package array

import "testing"

func TestIsContinuous(t *testing.T) {
	t.Log(isContinuous([]int{1, 2, 3, 0, 5}))
	t.Log(isContinuous([]int{1, 2, 3, 0, 8}))
}
