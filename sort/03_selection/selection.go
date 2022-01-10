package _3_selection

//寻找最小元素，与当前元素交换
func SelectSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

}
