package _1_bubble

import (
	mySort "data_structure/sort"
	"testing"
)

/*
单元测试
go test -v
-v 查看测试函数名称和运行时间

go test -v -cover
-cover 查看测试覆盖率

go test -v -cover -coverprofile bub.out
-coverprofile 覆盖率信息输出到一个文件
go tool cover -html=bub.out
浏览器窗口展示覆盖信息
-----------------------------
性能测试
go test
-bench 指定函数名（函数名就是BenchmarkBubSort1去掉Benchmark前缀），或者.（表示执行所有基准测试函数）。该字段是一个正则表达式。

输出结果的字段解释：
BenchmarkBubSort1-16：对BubSort1函数进行基准测试，16表示GOMAXPROCS，表示了并行能力
第二列的数字：被测函数执行的总次数
第三列xxx ns/op：每次操作的平均耗时

-benchmem 获得内存分配的统计数据
第四列xxx B/op：每次操作分配的内存大小
第五列xxx allocs/op：每次操作的内存分配次数
-----------------------------
*/

func TestBubble1(t *testing.T) {
	BubSort1(mySort.Sli1)

	isSort := mySort.CheckSlice(mySort.Sli1)
	if !isSort {
		t.Error(mySort.Sli1)
		return
	}

	if isSort {
		t.Logf("Bubble1 normal")
	}
	t.Log(mySort.Sli1)
}

func TestBubble2(t *testing.T) {
	BubSort2(mySort.Sli2)

	isSort := mySort.CheckSlice(mySort.Sli2)
	if !isSort {
		t.Error(mySort.Sli2)
		return
	}

	if isSort {
		t.Logf("Bubble2 normal")
	}

}

func TestBubble3(t *testing.T) {
	BubSort3(mySort.Sli3)

	isSort := mySort.CheckSlice(mySort.Sli3)
	if !isSort {
		t.Error(mySort.Sli3)
		return
	}

	if isSort {
		t.Logf("Bubble3 normal")
	}

}



func BenchmarkBubSort1(b *testing.B) {
	//默认上限时间是1s
	//先将b.N设置为1，执行性能测试函数BenchmarkBubSort1。执行时间没超过上限时间，就加大b.N，再执行该函数。直到执行时间大于等于上限时间为止。
	//此时，b.N的值就被包含在测试结果中，就是第二列的整数。
	//它表达的是，在最后一次测试中，BenchmarkBubSort1被执行了1次，总耗时大于等于了上限时间。BubSort1被执行了那个整数次。
	for i := 0; i < b.N; i++ {
		//b.StopTimer()
		////generate slice
		//sli := make([]int, 0, 10)
		//for j := 0; j < 10; j++ {
		//	sli = append(sli, rand.Intn(10))
		//}
		//b.StartTimer()
		BubSort1([]int{9, 1, 5, 8, 3, 7, 4, 6, 2})
	}
}

func BenchmarkBubSort1Num10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubSort1([]int{9, 1, 5, 8, 3, 10, 7, 4, 6, 2})
	}
}

func BenchmarkBubSort1Num20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubSort1([]int{9, 1, 5, 8, 3, 10, 7, 4, 6, 2, 20, 19, 18, 17, 16, 15, 14, 12, 11})
	}
}

//80ns/op
func BenchmarkBubSort3Num20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubSort3([]int{9, 1, 5, 8, 3, 10, 7, 4, 6, 2, 20, 19, 18, 17, 16, 15, 14, 12, 11})
	}
}

