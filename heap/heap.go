package heap

/*
参考标准库container/heap的实现，添加阅读注释。
up，down是核心方法，其余方法基本都是在调用他们。
*/

type MyHeap []int

//对应标准库中Less函数签名，改名为priority更准确些。当arr[i]应该排在arr[j]之前，函数返回true。
func (h *MyHeap) priority(i, j int) bool { return (*h)[i] < (*h)[j] }

//Init 对应建堆，从n/2-1 到 0，自上向下堆化
func (h *MyHeap) Init() {
	n := len(*h)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

//Push 对应新增元素，自下向上堆化
func (h *MyHeap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

//Pop 对应删除元素，自上向下堆化
func (h *MyHeap) Pop() (val int) {
	n := len(*h) - 1
	(*h)[0], (*h)[n] = (*h)[n], (*h)[0]
	h.down(0, n)

	*h, val = (*h)[:n], (*h)[n]
	return val
}

//func (h *MyHeap) Remove(i int) (val int) {
//	n := len(*h) - 1
//	if n != i {
//		(*h)[i], (*h)[n] = (*h)[n], (*h)[i]
//		if !h.down(i, n) {
//			h.up(i)
//		}
//	}
//	*h, val = (*h)[:n], (*h)[n]
//	return val
//}
//
//func (h *MyHeap) Fix(i int) {
//	if !h.down(i, len(*h)) {
//		h.up(i)
//	}
//}

//自下向上堆化
func (h *MyHeap) up(j int) {
	for {
		i := (j - 1) / 2 //父节点
		//子节点优先级低于父节点，priority返回false，取反为true，break ==》不再堆化
		//子节点优先级高于父节点，priority返回true，取反为false ==》继续堆化
		//【注意】：这里的索引越界判断，当j = 2, 1, 0时，i都等于0，也即i，j最终都会收敛为0，不会越界。
		if i == j || !h.priority(j, i) {
			break
		}

		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		j = i
	}
}

//自上向下堆化，n表示索引上限，不能等于该值。只要发生一次堆化，返回true。
func (h *MyHeap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		//索引越界，或整数溢出
		if j1 >= n || j1 < 0 {
			break
		}

		j := j1
		if j2 := j1 + 1; j2 < n && h.priority(j2, j1) {
			j = j2
		}
		//子节点优先级低于父节点，priority返回false，取反为true，break ==》不再堆化
		//子节点优先级高于父节点，priority返回true，取反为false ==》继续堆化
		if !h.priority(j, i) {
			break
		}
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		i = j
	}
	//只要发生一次堆化，返回true
	return i > i0
}
