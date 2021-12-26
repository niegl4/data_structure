package array

import "testing"

func TestAllOrder(t *testing.T) {
	t.Log(allOrder([]string{"a"}))
	t.Log(allOrder([]string{"a", "b"}))
	t.Log(allOrder([]string{"a", "b", "c"}))
}

func TestAllCombination(t *testing.T) {
	t.Log(allCombination([]string{"a"}))
	t.Log(allCombination([]string{"a", "b"}))
	t.Log(allCombination([]string{"a", "b", "c"}))
	t.Log(allCombination([]string{"a", "b", "c", "d"}))
}

func TestNumsOfQueens(t *testing.T) {
	res := numsOfQueens(8)
	for _, r := range res {
		t.Log(r)
	}
}
