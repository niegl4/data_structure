package array

import "testing"

func TestMergeSortArray(t *testing.T) {
	var (
		arr1 []int
		arr2 = []int{1, 2}
	)
	mergeSort(&arr1, arr2)
	t.Log(arr1)

	mergeSort(nil, arr2)

	arr1 = []int{1, 2}
	arr2 = []int{3, 4, 5, 6}
	mergeSort(&arr1, arr2)
	t.Log(arr1)

	arr2 = []int{1, 2}
	arr1 = []int{3, 4, 5, 6}
	mergeSort(&arr1, arr2)
	t.Log(arr1)
}
