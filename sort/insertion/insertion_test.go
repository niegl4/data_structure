package insertion

import (
	mySort "leecode/sort"
	"testing"
)

func TestInsertion(t *testing.T) {
	t.Log(mySort.Sli1)
	InsertionSort(mySort.Sli1)

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
