package _1_single_linked_list

/*
二十二-1
【获取】倒数第k个节点：

先获取正数第N个节点，再快、慢指针一起移动。（这种做法很好的平衡了N接近链表尾部/头部的情况。）
快慢指针的间距其实是k-1，这样，当快指针指向最后一个节点时，慢指针刚好指向倒数第k个节点。
*/
func getBottomK(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}
	fast := head
	i := 1 //fast提前步进k-1步
	for ; i <= k-1 && fast.next != nil; i++ {
		fast = fast.next
	}
	//k大于链表长度时，i=链表长度<k；k等于链表长度时，i=k=链表长度；k小于链表长度时，i=k
	if i < k {
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
【删除】倒数第k个节点：

v1相当于"获取倒数第k个节点"+"O(1)时间删除节点"的组合。
缺点是需要多次遍历，而且要实现两个子函数。
*/
func delBottomKV1(head **ListNode, k int) {
	bottomKNode := getBottomK(*head, k)
	if bottomKNode == nil {
		return
	}
	_ = delNode(head, bottomKNode)
	return
}

/*
v2不用多次遍历链表

快指针先指向正数第k个节点，
快慢指针一起移动，为了方便单链表删除，再新增一个慢指针的前继指针
*/
func delBottomKV2(head **ListNode, k int) {
	if head == nil || *head == nil || k <= 0 {
		return
	}
	fast := *head
	i := 1
	for ; i <= k-1 && fast.next != nil; i++ {
		fast = fast.next
	}
	if i < k {
		return
	}

	if fast.next == nil { //等价于删除头节点。eg：三个节点，删除倒数第3个。
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

/*
二十二-3
求链表的中间节点
如 1-2-3，返回2；1-2-3-4，返回2

元素个数为奇数：返回中间节点
元素个数为偶数：返回”上中位数“
*/
func midNodeOfList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
