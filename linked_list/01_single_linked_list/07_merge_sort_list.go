package _1_single_linked_list

/*
二十五
合并两个有序链表
*/
//空间复杂度O(1)，时间复杂度O(m+n)
func mergeSortList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var newHead, pre *ListNode
	for l1 != nil && l2 != nil {
		val1 := l1.value.(int)
		val2 := l2.value.(int)
		var node *ListNode
		if val2 < val1 {
			node = l2
			l2 = l2.next
		} else {
			node = l1
			l1 = l1.next
		}
		if pre == nil {
			newHead = node
			pre = node
		} else {
			pre.next = node
			pre = pre.next
		}
	}
	if l1 == nil {
		for l2 != nil && pre != nil {
			pre.next = l2
			l2 = l2.next
		}
	} else {
		for l1 != nil && pre != nil {
			pre.next = l1
			l1 = l1.next
		}
	}
	return newHead
}
