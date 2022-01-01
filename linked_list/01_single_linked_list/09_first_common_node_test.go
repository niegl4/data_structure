package _1_single_linked_list

import "testing"

func TestFirstCommonNode(t *testing.T) {
	list1 := genList([]int{1, 2, 3, 6, 7})
	list2 := genList([]int{4, 5})
	list2.next.next = list1.next.next.next
	node := firstCommonNode(list1, list2)
	t.Log(node.value)
}
