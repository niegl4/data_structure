package _7_shell

import (
	mySort "data_structure/sort"
	"testing"
)

func TestShellSort(t *testing.T) {
	t.Log(mySort.Sli1)
	ShellSort(mySort.Sli1)

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
