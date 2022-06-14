package array

import "testing"

func TestFindNumsWithSum(t *testing.T) {
	t.Log(findNumsWithSum([]int{1, 2, 4, 7, 11, 15}, 15))
}

func TestContinuousSeq(t *testing.T) {
	t.Log(continuousSeq(15))
	t.Log(continuousSeq(16))
}
