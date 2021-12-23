package _1_single_linked_list

/*
二十二-1
获取倒数第k个节点：

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
	if fast.next == nil && i < k {
		return nil
	}

	slow := head
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	return slow
}

/*
二十二-2
删除倒数第k个节点：

先获取正数第N个节点，再快、慢指针一起移动。这种做法很好的平衡了N接近链表尾部/头部的情况。
*/
//v1相当于"获取倒数第k个节点"+"O(1)时间删除节点"的组合。缺点是需要多次遍历，而且要实现两个子函数。
func delBottomKV1(head **ListNode, k int) {
	bottomKNode := getBottomK(*head, k)
	if bottomKNode == nil {
		return
	}
	_ = DelNode(head, bottomKNode)
	return
}

//v2不用多次遍历链表
func delBottomKV2(head **ListNode, k int) {
	if head == nil || *head == nil || k == 0 {
		return
	}
	fast := *head
	i := 1
	for ; i <= k-1 && fast.next != nil; i++ {
		fast = fast.next
	}
	//检测上方循环退出的原因。满足下面条件，说明list节点数小于k。
	if fast.next == nil && i < k {
		return
	}

	if fast.next == nil {//等价于删除头节点。eg：三个节点，删除倒数第3个。
		*head = (*head).next
		return
	}

	pre := *head
	slow := *head
	for fast.next != nil {
		fast = fast.next

		pre = slow
		slow = slow.next
	}
	pre.next = slow.next
	return
}
