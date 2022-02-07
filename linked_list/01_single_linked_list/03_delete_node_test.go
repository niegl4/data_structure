package _1_single_linked_list

import (
	"fmt"
	"testing"
)

func TestDelNode(t *testing.T) {
	testSet := [][]int{
		{1}, //删除头节点，剩余空链表
		{1, 2}, //删除头节点，剩余非空链表
		{1, 1}, //删除尾节点
		{1, 2, 2}, //删除中间节点
	}
	for idx, list := range testSet {
		headNode := genList(list)
		switch idx {
		case 0, 1:
			err := DelNode(&headNode, headNode)
			if err != nil {
				t.Fatal(idx, err)
			}
		case 2:
			err := DelNode(&headNode, headNode.next)
			if err != nil {
				t.Fatal(idx, err)
			}
		case 3:
			err := DelNode(&headNode, headNode.next)
			if err != nil {
				t.Fatal(idx, err)
			}
		}
		printList(headNode)
		fmt.Println("-------")
	}
}

func TestDelDupNode(t *testing.T) {
	testSet := [][]int{
		{1, 2}, //不包含重复元素
		{1, 1}, //头重复
		{1, 2, 2}, //尾重复
		{1, 2, 2, 3}, //中间重复
		{1, 1, 2, 3, 3, 4, 5, 5}, //头，中间，尾都有重复
	}
	for _, list := range testSet {
		headNode := genList(list)
		DelDupNode(&headNode)
		printList(headNode)
		fmt.Println("-------")
	}
}

func TestMidNodeOfList(t *testing.T) {
	l1 := genList([]int{1, 2, 3})
	l2 := genList([]int{1, 2, 3, 4})
	t.Log(midNodeOfList(l1).value)
	t.Log(midNodeOfList(l2).value)
}
