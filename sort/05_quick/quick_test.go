package _5_quick

import (
	mySort "data_structure/sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Log(mySort.Sli1)
	res := QuickSort(mySort.Sli1)

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


func TestQuickSort2(t *testing.T) {
	t.Log(mySort.Sli1)
	res := QuickSort2(mySort.Sli1)

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

func TestQuickSearch(t *testing.T) {
	//第1大为：12
	//第2大为：5
	//第3大为：4
	//第4大为：3
	//第5大为：2
	resMap := map[int]int {
		1: 12,
		2: 5,
		3: 4,
		4: 3,
		5: 2,
	}
	for i := 1; i <= 5; i++ {
		testArr := []int{4, 2, 5, 12, 3}
		ele := QuickSearch(testArr, i)
		if ele != resMap[i] {
			t.Fatal("search failed")
		}
		t.Log(ele)
	}
}