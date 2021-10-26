package _5_quick

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

//原地排序版快排，即空间复杂度是O(1)
func QuickSort2(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	quickSort2(arr, 0, n-1)
	return arr
}

//p:数组下标起点, r:数组下标终点
func quickSort2(arr []int, p, r int) {
	if p >= r {
		return
	}
	q := partition2(arr, p, r)
	quickSort2(arr, p, q-1)
	quickSort2(arr, q+1, r)
}

//空间复杂度O(1)，关键就在该分区函数中。比pivot小，才与索引i的元素交换。循环跳出时，索引比i小的元素，都比pivot小。
func partition2(arr []int, p, r int) (q int) {
	pivot := arr[r]
	//注意：i，j不同步。循环结束，i就是arr[r]的位置
	i := p
	for j := p; j <= r-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}
