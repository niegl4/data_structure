package _1_single_linked_list

import (
	"fmt"
	"testing"
)

func TestGetBottomK(t *testing.T) {
	head := genList([]int{1, 2, 3, 4})
	for i := 0; i < 6; i++ {
		bottomKNode := getBottomK(head, i)
		if bottomKNode == nil {
			t.Log("nil")
			continue
		}
		t.Log(bottomKNode.value)
	}
}

func TestDelBottomK(t *testing.T) {
	arrSet := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	for idx, arr := range arrSet {
		head := genList(arr)
		//delBottomKV1(&head, idx)
		delBottomKV2(&head, idx)
		printList(head)
		fmt.Println("-------")
	}
}

func TestMidNodeOfList(t *testing.T) {
	l1 := genList([]int{1, 2, 3})
	l2 := genList([]int{1, 2, 3, 4})
	t.Log(midNodeOfList(l1).value)
	t.Log(midNodeOfList(l2).value)
}
