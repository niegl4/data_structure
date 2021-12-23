package _1_single_linked_list

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	arrSet := [][]int{
		{1, 2, 3},
		nil,
		{1},
	}
	for _, arr := range arrSet {
		head := genList(arr)
		newHead := reverse(head)
		printList(newHead)
		fmt.Println("---------------------")
	}
}
