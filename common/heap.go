package common

type MaxHeap []int

func (m *MaxHeap) Len() int { return len(*m) }

func (m *MaxHeap) Less(i, j int) bool {
	return (*m)[i] >= (*m)[j]
}

func (m *MaxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
	return
}

func (m *MaxHeap) Push(data interface{}) {
	*m = append(*m, data.(int))
}

func (m *MaxHeap) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

type MinHeap []int

func (m *MinHeap) Len() int { return len(*m) }

func (m *MinHeap) Less(i, j int) bool {
	return (*m)[i] <= (*m)[j]
}

func (m *MinHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
	return
}

func (m *MinHeap) Push(data interface{}) {
	*m = append(*m, data.(int))
}

func (m *MinHeap) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}
