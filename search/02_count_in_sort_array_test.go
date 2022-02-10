package search

import "testing"

func TestCountInSortArray(t *testing.T) {
	t.Log(countInSortArray([]int{1, 2, 3, 3, 3, 3, 4, 5}, 3))
}

func TestMissNumInSortArray(t *testing.T) {
	t.Log(missNumInSortArray([]int{0, 1, 3, 4}))
}

func TestNumberEqualIndex(t *testing.T) {
	t.Log(numberEqualIndex([]int{-3, -1, 1, 3, 4}))
}
