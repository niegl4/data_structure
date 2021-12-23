package _1_single_linked_list

import "testing"

func TestMergeSortList(t *testing.T) {
	test1Arr1 := []int{1}
	test1Arr2 := []int{1}
	test2Arr1 := []int{1, 3, 3, 5, 7}
	test2Arr2 := []int{2, 4, 6, 8}
	test1L1 := genList(test1Arr1)
	test1L2 := genList(test1Arr2)
	test1Head := mergeSortList(test1L1, test1L2)
	printList(test1Head)

	test2L1 := genList(test2Arr1)
	test2l2 := genList(test2Arr2)
	test2Head := mergeSortList(test2L1, test2l2)
	printList(test2Head)
}
