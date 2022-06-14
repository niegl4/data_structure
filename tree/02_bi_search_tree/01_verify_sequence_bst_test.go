package _2_bi_search_tree

import "testing"

func TestVerifySequenceOfBstPostOrder(t *testing.T) {
	t.Log(verifySequenceOfBstPostOrder([]int{5, 7, 6, 9, 11, 10, 8}))
	t.Log(verifySequenceOfBstPostOrder([]int{7, 4, 6, 5}))
	t.Log(verifySequenceOfBstPostOrder([]int{2, 1}))
	t.Log(verifySequenceOfBstPostOrder([]int{1, 2}))
}

func TestVerifySequenceOfBstPreOrder(t *testing.T) {
	t.Log(verifySequenceOfBstPreOrder([]int{8, 6, 5, 7, 10, 9, 11}))
	t.Log(verifySequenceOfBstPreOrder([]int{5, 7, 4, 6}))
	t.Log(verifySequenceOfBstPreOrder([]int{2, 1}))
	t.Log(verifySequenceOfBstPreOrder([]int{1, 2}))
}
