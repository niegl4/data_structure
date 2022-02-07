package _1_single_linked_list

import "fmt"

// 单链表节点
type ListNode struct {
	next  *ListNode
	value interface{}
}

// 根据数据集合，生成单链表
func genList(arr []int) (head *ListNode) {
	var pre *ListNode
	for idx, num := range arr {
		node := &ListNode{value: num}
		if idx == 0 {
			head = node
		}
		if pre != nil {
			pre.next = node
			pre = node
		} else {
			pre = node
		}
	}
	return head
}

// 集中打印单链表
func printList(head *ListNode) {
	node := head
	result := make([]interface{}, 0)
	for node != nil {
		result = append(result, node.value)
		node = node.next
	}
	fmt.Println(result)
}
