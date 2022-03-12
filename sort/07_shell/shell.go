package _7_shell

func ShellSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	//gap缩小为1，是循环结束的条件.【gap就是分组的数目】
	for gap := length/2; gap > 0; gap /= 2 {

		//从每个分组的第二个元素，直到数组末尾
		for i := gap; i < length; i++ {

			//有点类似冒泡排序，每个分组中，把最小元素向前移动
			for j := i - gap; j >= 0; j -= gap {
				if arr[j+gap] < arr[j] {
					arr[j], arr[j+gap] = arr[j+gap], arr[j]
				} else {
					break
				}
			}

		}

	}
}
