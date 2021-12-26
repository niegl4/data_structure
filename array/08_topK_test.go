package array

import "testing"

func TestTopK(t *testing.T) {
	arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	for i := 0; i < 15; i++ {
		t.Logf("tpk %d--------", i)
		num, err := topK(arr, i)
		if err != nil {
			t.Log(err)
		} else {
			t.Log(num)
		}
		t.Log("--------")
	}
}
