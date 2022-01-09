package _6_heap

import (
	mySort "data_structure/sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
	t.Log(mySort.Sli1)
	HeapSort(mySort.Sli1)

	isSort := mySort.CheckSlice(mySort.Sli1)
	if !isSort {
		t.Error(mySort.Sli1)
		return
	}

	if isSort {
		t.Log("normal")
	}
	t.Log(mySort.Sli1)
}
