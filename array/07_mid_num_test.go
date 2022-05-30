package array

import "testing"

func TestMidNumByPartition(t *testing.T) {
	arrSet := [][]int{
		{1, 2, 3, 2, 2, 2, 5, 4, 2},
		{1, 2, 2, 4, 5},
		{1, 2, 2, 2, 5},
		{1, 1},
	}
	for _, arr := range arrSet {
		num, err := midNumByPartition(arr)
		if err != nil {
			t.Log(err)
		} else {
			t.Log(num)
		}
		t.Log("-----------")
	}
}

func TestMidNumBySkill(t *testing.T) {
	arrSet := [][]int{
		{1, 2, 3, 2, 2, 2, 5, 4, 2},
		{1, 2, 2, 4, 5},
		{1, 1},
	}
	for _, arr := range arrSet {
		num, err := midNumBySkill(arr)
		if err != nil {
			t.Log(err)
		} else {
			t.Log(num)
		}
		t.Log("-----------")
	}
}
