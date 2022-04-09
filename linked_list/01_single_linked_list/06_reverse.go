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
		cur       = head
	)
	for cur.next != nil {
		//临时保存next
		next = cur.next

		//修改节点指针指向
		cur.next = pre

		//pre，node步进
		pre = cur
		cur = next
	}
	cur.next = pre
	return cur
}
