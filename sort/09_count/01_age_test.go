package _9_count

import "testing"

func TestSortAges(t *testing.T) {
	ages := []int{1, 1, 99, 99, 4, 6, 5}
	err := sortAge(ages)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ages)
}
