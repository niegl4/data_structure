package _6_heap

/*
参考标准库sort中堆排序的实现，添加阅读注释。
*/

//HeapSort 从小到大排序，建立的是大顶堆
func HeapSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	//建堆：自顶向下堆化
	for i := (length - 1) / 2; i >= 0; i-- {
		siftDown(arr, i, length)
	}

	//逐次弹出堆顶元素（0号索引），弹出放在队尾，这样最终就是从小到大排列。
	//弹出堆顶元素：自顶向下堆化
	for i := length - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		siftDown(arr, 0, i)
	}
}

//siftDown其实就是自上向下堆化，lo：要堆化元素的索引；hi：堆的长度。
func siftDown(arr []int, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && arr[child] < arr[child+1] {
			child++
		}

		if arr[root] >= arr[child] {
			return
		}
		arr[root], arr[child] = arr[child], arr[root]
		root = child
	}
}
