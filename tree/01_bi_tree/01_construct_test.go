package _1_bi_tree

import (
	"testing"
)

func TestConstructByPreAndIn(t *testing.T) {
	root, err := constructByPreAndIn([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	if err != nil {
		t.Fatal(err)
	}

	tmp := root.PreOrderTraverse()
	t.Log(tmp)

	tmp = root.InOrderTraverse()
	t.Log(tmp)

	tmp = root.PostOrderTraverse()
	t.Log(tmp)

	levelRes := root.LevelOrderTraverse()
	t.Log(levelRes)
}

func TestConstructByPostAndIn(t *testing.T) {
	root, err := constructByPostAndIn([]int{7, 4, 2, 5, 8, 6, 3, 1}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	if err != nil {
		t.Fatal(err)
	}

	tmp := root.PreOrderTraverse()
	t.Log(tmp)

	tmp = root.InOrderTraverse()
	t.Log(tmp)

	tmp = root.PostOrderTraverse()
	t.Log(tmp)

	levelRes := root.LevelOrderTraverse()
	t.Log(levelRes)
}
