package array

import "testing"

func TestNumAppearOnce(t *testing.T) {
	t.Log(twoNumAppearOnce([]int{2, 4, 3, 6, 3, 2, 5, 5}))
}

func TestOneNumAppearOnce(t *testing.T) {
	t.Log(oneNumAppearOnce([]int{3, 3, 3, 2, 2, 2, 1}))
}
