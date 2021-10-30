package _2_insertion

import (
	mySort "data_structure/sort"
	"testing"
)

func TestInsertion(t *testing.T) {
	t.Log(mySort.Sli1)
	InsertSort(mySort.Sli1)

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

//30ns/op，同样的测试数据与冒泡排序（80ns/op）对比，插入排序确实比冒泡排序优越
func BenchmarkInsertSortNum20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort([]int{9, 1, 5, 8, 3, 10, 7, 4, 6, 2, 20, 19, 18, 17, 16, 15, 14, 12, 11})
	}
}
