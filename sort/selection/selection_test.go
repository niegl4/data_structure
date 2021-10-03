package selection

import (
	mySort "data_structure/sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	t.Log(mySort.Sli1)
	SelectSort(mySort.Sli1)

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


