package _1_single_linked_list

import "testing"

func TestGetBottomK(t *testing.T) {
	head := genList([]int{1, 2, 3, 4})
	for i := 0; i < 6; i++ {
		bottomKNode := getBottomK(head, i+10)
		if bottomKNode == nil {
			t.Log("nil")
			continue
		}
		t.Log(bottomKNode.value)
	}
}
