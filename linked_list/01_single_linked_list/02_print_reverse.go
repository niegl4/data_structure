package _1_single_linked_list

import (
	"data_structure/common"
	"fmt"
)

/*
六
单链表从尾到头打印，不允许修改链表

单链表从尾到头打印:借助栈实现
时间复杂度O(n)，空间复杂度O(n)
*/
func (n *ListNode) printReverse() {
	if n == nil {
		return
	}
	stack := common.NewStack()
	next := n
	for next != nil {
		stack.Push(next.value.(int))
		next = next.next
	}
	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}
}

func (n *ListNode) printReverseV2() {
	if n == nil {
		return
	}
	stack := make([]interface{}, 0, 8)
	next := n
	for next != nil {
		stack = append(stack, next)
		next = next.next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Println(stack[i])
	}
}
