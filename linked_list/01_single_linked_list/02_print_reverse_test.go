package _1_single_linked_list

import "testing"

func TestListNode_PrintReverse(t *testing.T) {
	headNode := genList([]int{1, 2, 3, 4, 5})

	headNode.printReverse()
	headNode.printReverseV2()
}
