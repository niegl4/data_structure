package _1_single_linked_list

import (
	"fmt"
	"testing"
)

func TestDelNode(t *testing.T) {
	testSet := [][]int{
		{1},
		{1, 2},
		{1, 1},
		{1, 2, 2},
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
		{1, 2},
		{1, 1, 2, 3, 3, 4, 5, 5},
		{1, 1},
		{1, 2, 2},
		{1, 2, 2, 3},
	}
	for _, list := range testSet {
		headNode := genList(list)
		//printList(headNode)
		//node1.DelDupNode()
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
