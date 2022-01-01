package _1_single_linked_list

/*
五十二
两个单链表的第一个公共节点。
*/
func firstCommonNode(list1, list2 *ListNode) *ListNode {
	list1Len := 0
	list2Len := 0
	node1 := list1
	node2 := list2
	for node1 != nil {
		list1Len++
		node1 = node1.next
	}
	for node2 != nil {
		list2Len++
		node2 = node2.next
	}
	if list1Len == 0 || list2Len == 0 {
		return nil
	}

	node1 = list1
	node2 = list2
	if list2Len > list1Len {
		dif := list2Len - list1Len
		for ; dif > 0; dif-- {
			node2 = node2.next
		}
	} else {
		dif := list1Len - list2Len
		for ; dif > 0; dif-- {
			node1 = node1.next
		}
	}

	for node1 != nil && node2 != nil {
		if node1 == node2 {
			return node1
		} else {
			node1 = node1.next
			node2 = node2.next
		}
	}
	return nil
}
