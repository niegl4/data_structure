package search

import "testing"

func TestSearchMinInRotateSortArr2(t *testing.T) {
	t.Log(searchMinInRotateSortArr2([]int{3, 4, 5, 1, 2}))
	t.Log(searchMinInRotateSortArr2([]int{1, 2, 3, 4, 5}))
	t.Log(searchMinInRotateSortArr2([]int{1, 0, 1, 1, 1}))
	t.Log(searchMinInRotateSortArr2([]int{1, 1, 1, 0, 1}))
}
