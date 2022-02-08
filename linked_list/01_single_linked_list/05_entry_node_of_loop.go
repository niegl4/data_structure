package _1_single_linked_list

/*
二十三
找到链表中环的入口节点。

1.如果有环，找到相遇的节点
2.通过相遇节点，计算环中节点的个数n
3.通过个数，找到环的入口节点
	寻找环的入口，n个节点就要先走n步（倒数第k个节点的问题中，最好是先走k-1步）
*/
func meetNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head.next
	for fast != nil && fast.next != nil {
		if slow == fast {
			return slow
		}
		slow = slow.next
		fast = fast.next.next
	}
	return nil
}

func entryNodeOfLoop(head *ListNode) *ListNode {
	meetingNode := meetNode(head)
	if meetingNode == nil {
		return nil
	}

	n := 1
	node := meetingNode.next
	for node != meetingNode {
		node = node.next
		n++
	}

	fast := head
	for i := 0; i < n; i++ {
		fast = fast.next
	}
	slow := head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	return slow
}
