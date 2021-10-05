package _5_quicksort

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
func Sort2(inputArr []int) []int {
	n := len(inputArr)
	if n < 2 {
		return inputArr
	}
	sort2(inputArr, 0, n-1)
	return inputArr
}

func sort2(inputArr []int, p, r int) {
	if p >= r {
		return
	}
	q := partition2(inputArr, p, r)
	sort2(inputArr, p, q-1)
	sort2(inputArr, q+1, r)
}

//空间复杂度O(1)，关键就在该分区函数中。比pivot小，才与索引i的元素交换。循环跳出时，索引比i小的元素，都比pivot小。
func partition2(inputArr []int, p, r int) (q int) {
	pivot := inputArr[r]
	i := p
	for j := p; j <= r-1; j++ {
		if inputArr[j] < pivot {
			inputArr[i], inputArr[j] = inputArr[j], inputArr[i]
			i++
		}
	}
	inputArr[i], inputArr[r] = inputArr[r], inputArr[i]
	return i
}
