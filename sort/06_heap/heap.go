package _6_heap

/*
参考标准库sort中堆排序的实现，添加阅读注释。
*/

/*
数组从索引0开始，建立堆。已知某节点索引为i：
1.左孩子：2*i+1。右孩子：2*i+2。
2.父节点：（i-1）/2。
3.建堆时，最后一个非叶节点为：length/2 - 1。
*/

//HeapSort 从小到大排序，建立的是大顶堆
func HeapSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	//todo：最后一个非叶节点应该是length/2 - 1,但是源码中是（length-1）/2。
	//建堆：自顶向下堆化
	for i := length/2 - 1; i >= 0; i-- {
		siftDown(arr, i, length)
	}

	//逐次弹出堆顶元素（0号索引），弹出放在队尾，这样最终就是从小到大排列。
	//弹出堆顶元素：自顶向下堆化
	for i := length - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		siftDown(arr, 0, i)
	}
}

//siftDown其实就是自上向下堆化，lo：要堆化元素的索引；hi：【注意：索引上限小于hi】。
func siftDown(arr []int, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			return
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
