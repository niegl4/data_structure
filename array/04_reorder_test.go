package array

import "testing"

func TestReorder(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	reorder(arr1, isOdd)
	t.Log(arr1)

	arr2 := []int{-1, 2, -3, -4, 5}
	reorder(arr2, isNegative)
	t.Log(arr2)

	arr3 := []int{-1, 2, -3, -4, 6}
	reorder(arr3, isDivisibleBy3)
	t.Log(arr3)
}
