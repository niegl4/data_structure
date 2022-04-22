package string

import "testing"

func TestLengthOfLongestSubStr(t *testing.T) {
	t.Log(lengthOfLongestSubStr("arabcacfr"))
}

func TestLongestIndex(t *testing.T) {
	t.Log(longestIndex([]int{0, 0, 0}, check1))
	t.Log(longestIndex([]int{1, 0, 0}, check1))
	t.Log(longestIndex([]int{0, 1, 0}, check1))
	t.Log(longestIndex([]int{0, 0, 1}, check1))
	t.Log(longestIndex([]int{0, 1, 1, 0, 1, 1, 1}, check1))
	t.Log(" ")
	t.Log(longestIndex([]int{0, 0, 0}, check0))
	t.Log(longestIndex([]int{1, 0, 0}, check0))
	t.Log(longestIndex([]int{0, 1, 0}, check0))
	t.Log(longestIndex([]int{0, 0, 1}, check0))
	t.Log(longestIndex([]int{0, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0}, check0))
}
