package quicksort

//边切分，边确保大致有序（左边小于中间元素，右边大于中间元素）
//最后从最小的单位逐级返回，每次返回都是有序的
func QuickSort(inputArr []int) []int {
	if len(inputArr) < 2 {
		return inputArr
	}
	left := make([]int, 0)  // var left []int
	right := make([]int, 0) // var left = []int{}

	midIndex := len(inputArr) / 2
	mid := inputArr[midIndex]
	//这步删除中间元素一定要做.否则可能造成无穷递归.
	//1, 2       9, 5, 8, 3, 7, 4, 6
	//空         9, 5, 8, 3, 7, 4, 6
	inputArr = append(inputArr[:midIndex], inputArr[midIndex+1:]...)
	for _, item := range inputArr {
		if item < mid {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}

	ret1 := QuickSort(left)
	ret2 := QuickSort(right)
	ret1 = append(ret1, mid)
	ret1 = append(ret1, ret2...)
	return ret1
}
