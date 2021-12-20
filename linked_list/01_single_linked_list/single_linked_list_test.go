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
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func TestListNode_PrintReverse(t *testing.T) {
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	node1.PrintReverse()
	//l3.PrintReverse()
}

func TestListNode_DelNode(t *testing.T) {
	headNode := genList([]int{1, 2, 3, 4, 5})
	//printList(headNode)
	err := headNode.DelNode(headNode)
	if err != nil {
		t.Fatal(err)
	}
	printList(headNode)

	headNode = genList([]int{1})
	//printList(headNode)
	err = headNode.DelNode(headNode)
	if err != nil {
		t.Fatal(err)
	}
	printList(headNode)

}

func TestListNode_DelDupNode(t *testing.T) {
	headNode := genList([]int{1, 1, 2, 3, 3, 4, 5, 5})
	printList(headNode)
	//node1.DelDupNode()
	DelDupNode(&headNode)
	printList(headNode)
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



func TestListNode_Reverse(t *testing.T) {
	t.Log(l1.Reverse().format())

	t.Log(l2.Reverse().format())

	t.Log(l3.Reverse().format())
}

func Test_HasCycle(t *testing.T) {
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node2
	l := &LinkedList{head: &ListNode{next: node1}}
	t.Log(l.HasCycle())
	t.Log(node1.HasCycle())

	node5.next = nil
	t.Log(l.HasCycle())
	t.Log(node1.HasCycle())
}


