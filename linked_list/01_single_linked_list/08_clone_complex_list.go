package _1_single_linked_list

/*
三十五
复制复杂链表。复杂链表中，除了next指向下一节点，还有sibling指针指向任意节点或指向nil。

最容易想到的办法是，先复制主干，再逐个节点处理sibling指针（每个sibling都是从头遍历查找），时间复杂度O(n^2)
借助散列表，在O(n)的空间复杂度下，时间复杂度被优化为O(n)。（原链表第n个节点的指针 与 复制链表第n个节点的指针 的映射）
下面的方法，不使用散列表，时间复杂度为O(n)
*/

type ComplexList struct {
	Val     interface{}
	Next    *ComplexList
	Sibling *ComplexList
}

//1.每个原链表的节点后面，都插入一个copy节点
//2.copy节点的sibling，指向它前一个节点的sibling的next
//3.把copy节点从原链表中断开
func cloneComplexList(head *ComplexList) *ComplexList {
	if head == nil {
		return nil
	}
	cloneNodes(head)
	connectSibling(head)
	return reconnectNodes(head)
}

func cloneNodes(head *ComplexList) {
	node := head
	for node != nil {
		next := node.Next
		node.Next = &ComplexList{
			Val:  node.Val,
			Next: next,
		}
		node = next
	}
}

func connectSibling(head *ComplexList) {
	node := head
	for node != nil {
		if node.Sibling == nil {
			node.Next.Sibling = nil
		} else {
			node.Next.Sibling = node.Sibling.Next
		}
		node = node.Next.Next
	}
}

func reconnectNodes(head *ComplexList) *ComplexList {
	cloneHead := head.Next
	node := head
	cloneNode := cloneHead
	for node != nil {
		next := cloneNode.Next
		if next == nil {
			node.Next = nil
			cloneNode.Next = nil
			break
		}
		node.Next = next
		cloneNode.Next = next.Next

		node = node.Next
		cloneNode = cloneNode.Next
	}
	return cloneHead
}
