package _1_single_linked_list

/*
删除倒数第k个节点：

先获取正数第N个节点，再快、慢指针一起移动。这种做法很好的平衡了N接近链表尾部/头部的情况。
 */
func getBottomK(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}
	fast := head
	i := 1
	for ; i <= k-1 && fast.next != nil; i++ {
		fast = fast.next
	}
	//检测上方循环退出的原因。满足下面条件，说明list节点数小于k。
	if fast.next == nil && i < k-1 {
		return nil
	}

	slow := head
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	return slow
}

func delBottomK(head **ListNode, k int) {
	if head == nil || *head == nil || k == 0 {
		return
	}
	fast := *head
	for i := 1; i <= k && fast.next != nil; i++ {
		fast = fast.next
	}
	if fast == nil {
		return
	}

	slow := *head
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	slow.next = slow.next.next
	return
}

func (l *LinkedList) DeleteBottomK(n int) {
	if n <= 0 || l == nil || l.head == nil || l.head.next == nil {
		return
	}

	//先获取正数第N个节点
	fast := l.head //从head开始，那么结束的时候，fast刚好就是第n个节点
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}
	//上面循环结束的原因，如果是fast为nil，而不是计数到达，那么说明链表内没有倒数第N个节点
	if fast == nil {
		return
	}


	slow := l.head
	for fast.next != nil {//fast为nil的时候，slow刚好是倒数第N+1个节点，方便删除。
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}
