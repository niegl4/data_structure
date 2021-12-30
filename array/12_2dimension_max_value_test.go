package array

import "testing"

func TestMaxValue(t *testing.T) {
	matrix := [][]int{
		{1, 10, 3, 8},
		{12, 2, 9, 6},
		{5, 7, 4, 11},
		{3, 7, 16, 5},
	}
	t.Log(maxValue(matrix))
}
