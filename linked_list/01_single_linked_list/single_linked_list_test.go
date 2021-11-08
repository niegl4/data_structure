package _1_single_linked_list

import "testing"

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
	)

	t.Log(l1.Reverse().format())

	t.Log(l2.Reverse().format())

	t.Log(l3.Reverse().format())
}
