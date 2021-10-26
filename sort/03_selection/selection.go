package _3_selection

func NoUse() {}

//寻找最小元素，与当前元素交换
func SelectSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

}
