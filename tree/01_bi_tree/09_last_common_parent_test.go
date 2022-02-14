package _1_bi_tree

import "testing"

func TestGetLastCommonParent(t *testing.T) {
	root, err := constructByPreAndIn([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	if err != nil {
		t.Fatal(err)
	}

	node := getLastCommonParent(root, root.lChild.lChild.rChild, root.rChild.rChild.lChild)
	t.Log(node.Data)
}
