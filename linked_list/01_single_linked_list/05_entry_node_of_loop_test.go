package _1_single_linked_list

import "testing"

func TestEntryNodeOfLoop(t *testing.T) {
	list1 := genList([]int{1})
	list1.next = list1
	node := entryNodeOfLoop(list1)
	if node != nil {
		t.Log(node.value)
	} else {
		t.Log("nil")
	}

	list2 := genList([]int{1, 2, 3})
	node = entryNodeOfLoop(list2)
	if node != nil {
		t.Log(node.value)
	} else {
		t.Log("nil")
	}

	list3 := genList([]int{1, 2, 3, 4, 5})
	list3.next.next.next.next = list3.next.next
	node = entryNodeOfLoop(list3)
	if node != nil {
		t.Log(node.value)
	} else {
		t.Log("nil")
	}
}
