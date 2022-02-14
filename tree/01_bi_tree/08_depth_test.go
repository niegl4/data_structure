package _1_bi_tree

import "testing"

func TestDepth(t *testing.T) {
	root, err := constructByPreAndIn([]int{1, 2, 4, 5, 7, 3, 6}, []int{4, 2, 7, 5, 1, 3, 6})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(depth(root))
}
