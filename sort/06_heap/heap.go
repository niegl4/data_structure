package _6_heap

//做升序排序，建大顶堆。从最后一个有子节点的节点开始，自上向下堆化。这里从1号索引开始存数据，n表示一共几个数据。
func HeapSort(a []int, n int) []int {
	buildHeap(a, n)

	k := n
	for k >= 1 {
		swap(a, 1, k)
		heapifyUpToDown(a, 1, k-1)
		k--
	}
	return a
}

//从最后一个有子节点的节点开始，自上向下堆化。
func buildHeap(a []int, n int) {
	//heapify from the last parent node
	for i := n / 2; i >= 1; i-- {
		heapifyUpToDown(a, i, n)
	}
}

//heapify from up to down , node index = top
func heapifyUpToDown(a []int, top int, count int) {

	for i := top; i <= count/2; { //必须小于等于count/2，大于则节点没有子节点

		maxIndex := i
		if a[i] < a[i*2] { //与左子节点比较
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] { //与右子节点比较
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		a[i], a[maxIndex] = a[maxIndex], a[i]
		i = maxIndex
	}

}

//swap two elements
func swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}