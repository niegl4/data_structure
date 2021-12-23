package _1_single_linked_list

import (
	"fmt"
	"testing"
)

var (
	l1 *ListNode
	l2 = &ListNode{value: 0, next: nil}
	l3 = &ListNode{
		value: 0,
		next: &ListNode{
			value: 1,
			next: &ListNode{
				value: 2,
				next: &ListNode{
					value: 3,
					next: &ListNode{
						value: 4,
						next:  nil,
					},
				},
			},
		},
	}
	node1 = &ListNode{value: 1, next: nil}
	node2 = &ListNode{value: 2, next: nil}
	node3 = &ListNode{value: 3, next: nil}
	node4 = &ListNode{value: 4, next: nil}
	node5 = &ListNode{value: 5, next: nil}
)

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

func printList(head *ListNode) {
	node := head
	result := make([]interface{}, 0)
	for node != nil {
		result = append(result, node.value)
		node = node.next
	}
	fmt.Println(result)
}

func TestLinkedList_Reverse(t *testing.T) {
	var (
		l1 *LinkedList
		l2 = &LinkedList{head: nil}
		l3 = &LinkedList{
			head: &ListNode{next: nil},
		}
		l4 = LinkedList{
			head: &ListNode{
				next: &ListNode{
					value: 1,
					next: &ListNode{
						value: 2,
						next: &ListNode{
							value: 3,
							next: &ListNode{
								value: 4,
								next:  nil,
							},
						},
					},
				},
			},
		}
	)
	l1.Reverse()
	t.Log(l1.format())

	l2.Reverse()
	t.Log(l2.format())

	l3.Reverse()
	t.Log(l3.format())

	l4.Reverse()
	t.Log(l4.format())
}


