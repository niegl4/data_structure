package _1_bi_tree

import "testing"

func TestFindPath(t *testing.T) {
	root, err := construct([]int{10, 5, 4, 7, 12}, []int{4, 5, 7, 10, 12})
	if err != nil {
		t.Fatal(err)
	}
	paths := findPath(root, 22)
	for _, path := range paths {
		var tmp []interface{}
		for _, n := range path {
			tmp = append(tmp, n.Data)
		}
		t.Log(tmp)
	}
}
