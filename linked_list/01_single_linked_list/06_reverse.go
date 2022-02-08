package _1_single_linked_list

/*
二十四
反转单链表

单链表反转其实只涉及一个操作：修改节点指针。仍然是pre，node，next配合。
*/

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.next == nil {
		return head
	}

	var (
		pre, next *ListNode
		node      = head
	)
	for node.next != nil {
		//临时保存next
		next = node.next

		//修改节点指针指向
		node.next = pre

		//pre，node步进
		pre = node
		node = next
	}
	node.next = pre
	return node
}
