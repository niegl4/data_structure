package array

import (
	"container/heap"
	"data_structure/common"
	"fmt"
)

/*
四十一
数据流中的中位数。
如果当前是奇数，排序后中间的数就是中位数。
如果当前是偶数，排序后中间的两个数的平均数就是中位数。
*/

func calMidNum(numberStream <-chan int) {
	var (
		maxH common.MaxHeap
		minH common.MinHeap
	)
	i := 1
	for {
		num := <-numberStream
		if i&0b1 == 1 {
			midNum := 0
			if maxH.Len() == 0 {
				midNum = num
				maxH.Push(num)
			} else {
				min := heap.Pop(&minH).(int)
				if num > min {
					heap.Push(&minH, num)
					heap.Push(&maxH, min)
				} else {
					heap.Push(&minH, min)
					heap.Push(&maxH, num)
				}
				midNum = heap.Pop(&maxH).(int)
				heap.Push(&maxH, midNum)
			}
			fmt.Printf("idx: %d, midNum:%d\n", i, midNum)
		} else {
			num1 := 0
			num2 := 0
			if minH.Len() == 0 {
				minH.Push(num)
				num1 = heap.Pop(&maxH).(int)
				heap.Push(&maxH, num1)
				num2 = num
			} else {
				max := heap.Pop(&maxH).(int)
				if num < max {
					heap.Push(&maxH, num)
					heap.Push(&minH, max)
				} else {
					heap.Push(&maxH, max)
					heap.Push(&minH, num)
				}
				num1 = heap.Pop(&maxH).(int)
				heap.Push(&maxH, num1)
				num2 = heap.Pop(&minH).(int)
				heap.Push(&minH, num2)
			}
			fmt.Printf("idx: %d, midNum:%d\n", i, (num1+num2)/2)
		}
		if i == 7 {
			break
		}
		i++
	}
}
