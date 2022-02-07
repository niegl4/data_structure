package _1_single_linked_list

import "errors"

/*
十八-1
在O(1)时间内删除单链表节点

关键是要求O(1)，所以通过“赋值”实现删除。
1.要删除的是头节点
2.要删除的不是尾节点
3.要删除的是尾结点
*/
func DelNode(head **ListNode, node *ListNode) error {
	if head == nil || *head == nil {
		return errors.New("list is nil")
	}

	if *head == node { //要删除节点是头节点
		*head = (*head).next
		return nil
	}
	if node.next != nil { //要删除节点不是尾节点
		node.value = node.next.value
		node.next = node.next.next
		return nil
	} else { //要删除节点是尾节点
		pre := *head
		for pre.next != nil && pre.next != node {
			pre = pre.next
		}
		if pre.next == nil {
			return errors.New("not in list")
		} else {
			pre.next = node.next
			return nil
		}
	}
}

/*
十八-2
排序单链表中，删除重复节点
如 1-2-3-3-4-4-5，删除后：1-2-5

pre,node,next三个指针配合。
pre：只有node与next不等时，pre才步进。也即，node与next还没判断，pre并不步进。
另外，还要小心连续相等区间出现在链表头部，此时需要更换头节点。
*/
func DelDupNode(head **ListNode) {
	if head == nil || *head == nil { //空链表
		return
	}

	if (*head).next == nil { //只有一个节点，必不可能重复
		return
	}

	var (
		preNode *ListNode
		node    = *head
	)
	for node != nil {
		nextNode := node.next
		if nextNode == nil { //当前是链表最后一个节点(最后一个元素也是重复元素，for循环条件已经排除这种可能)
			break
		}
		if node.value != nextNode.value { //node节点与next节点值不等：移动pre，node
			preNode = node
			node = nextNode
			continue
		}

		//连续多个相等
		val := node.value
		toDelNode := node
		for toDelNode != nil && toDelNode.value == val { //循环退出时：toDelNode指向离开连续相等区间的第一个节点
			toDelNode = toDelNode.next
		}

		//连续区间出现在开头
		if node == *head { //从头开始连续相等：链表头节点更换
			node = toDelNode
			*head = toDelNode
		} else { //不是从头开始连续相等
			if preNode != nil {
				preNode.next = toDelNode
				node = toDelNode
			} else {
				node = toDelNode
			}
		}
	}
}

/*
十八-3
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
