package _3_selection

func NoUse() {}

//寻找最小元素，与当前元素交换
func SelectSort(input []int) {
	n := len(input)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if input[min] > input[j] {
				min = j
			}
		}

		if min != i {
			input[i], input[min] = input[min], input[i]
		}
	}

}
