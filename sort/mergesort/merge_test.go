package mergesort

import (
	mySort "leecode/sort"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Log(mySort.Sli1)
	res := MergeSort(mySort.Sli1)

	isSort := mySort.CheckSlice(res)
	if !isSort {
		t.Error(res)
		return
	}

	if isSort {
		t.Logf("normal")
	}
	t.Log(res)
}
