package _1_single_linked_list

//单链表节点
type ListNode struct {
	next  *ListNode
	value interface{}
}

//单链表。在任何时候，不管链表是不是空，head 指针都会一直指向一个哨兵结点，哨兵节点再指向链表的第一个节点。
type LinkedList struct {
	head *ListNode
}

//单链表反转：pre，cur，next三个指针的步进
func (l *LinkedList) Reverse() {
	//l为空 or l的哨兵节点字段为空 or 空链表 or 只有一个节点的链表
	if l == nil || l.head == nil || l.head.next == nil || l.head.next.next == nil {
		return
	}
	//pre，cur，next三个临时变量，从cur=第一个链表节点开始遍历
	pre := new(ListNode)
	cur := l.head.next
	for cur != nil {
		//获取next，因为马上就要对cur.next修改，而next就是下一个循环周期中的cur
		next := cur.next

		//单链表的反转操作，其实就是这一步
		cur.next = pre

		//pre，cur的步进
		pre = cur
		cur = next
	}
	//l.head.next就是链表的第一个元素。循环结束的时候，pre就是原来链表的最后一个元素。
	l.head.next = pre
}

//判断单链表是否有环：快指针，步长为2；慢指针，步长为1；如果有环，它们一定会相遇。
func (l *LinkedList) HasCycle() bool {
	if l == nil || l.head == nil {
		return false
	}

	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}