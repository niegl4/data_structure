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
func (n *ListNode) PrintReverse() {
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
