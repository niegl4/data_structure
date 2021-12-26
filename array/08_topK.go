package array

import (
	"container/heap"
	"errors"
)

/*
四十
在数组或者数字流中，找到最小的k个数。

最小的k个数 等价于 从小到大排序的topK。
基于partition函数实现，会修改原数组，时间复杂度为O(n)
基于大顶堆实现，不修改原数组，时间复杂度O(nlogn)
*/
func topK(arr []int, k int) ([]int, error) {
	if k < 0 {
		return nil, errors.New("invalid k")
	}
	if k == 0 {
		return nil, nil
	}

	length := len(arr)
	if length < 1 || length <= k {
		return arr, nil
	}

	var maxH maxHeap
	for idx, num := range arr {
		if idx < k {
			maxH.Push(num)
		} else {
			if idx == k {
				heap.Init(&maxH)
			}
			curMax := heap.Pop(&maxH).(int)
			if num < curMax {
				heap.Push(&maxH, num)
			} else {
				heap.Push(&maxH, curMax)
			}
		}
	}
	return maxH, nil
}

type maxHeap []int

func (m *maxHeap) Len() int { return len(*m) }

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i] >= (*m)[j]
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
	return
}

func (m *maxHeap) Push(data interface{}) {
	*m = append(*m, data.(int))
}

func (m *maxHeap) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}
