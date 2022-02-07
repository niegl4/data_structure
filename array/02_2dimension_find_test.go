package array

import "testing"

func TestFind(t *testing.T) {
	matrix := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	t.Log(find(matrix, 7))
	t.Log(find(matrix, 5))
}
