package _1_single_linked_list

/*
二十四
反转单链表

单链表反转其实只涉及一个操作：修改节点指针。仍然是pre，node，next配合。
*/

func reverse(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}

	var (
		cur       = head
		pre, next *ListNode
	)
	for cur != nil {
		next = cur.next

		cur.next = pre
		pre = cur
		cur = next
	}
	return pre
}
