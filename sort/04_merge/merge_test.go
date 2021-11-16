package _4_merge

import (
	mySort "data_structure/sort"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Log(mySort.Sli1)
	MergeSort(mySort.Sli1, 9)

	isSort := mySort.CheckSlice(mySort.Sli1)
	if !isSort {
		t.Error(mySort.Sli1)
		return
	}

	if isSort {
		t.Logf("normal")
	}
	t.Log(mySort.Sli1)
}
