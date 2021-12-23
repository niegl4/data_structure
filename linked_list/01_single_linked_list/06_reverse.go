package _1_single_linked_list

/*
二十四
反转单链表
 */

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.next == nil {
		return head
	}

	var (
		pre, next  *ListNode
		node = head
	)
	for node.next != nil {
		next = node.next
		node.next = pre

		pre = node
		node = next
	}
	node.next = pre
	return node
}
