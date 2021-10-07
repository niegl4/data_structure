package _4_merge

//先切分，再排序
//最后从最小的单位逐级返回，每次返回都是有序的
func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	midIndex := len(arr) / 2
	left := arr[:midIndex]
	right := arr[midIndex:]
	arr1 := MergeSort(left)
	arr2 := MergeSort(right)
	return merge(arr1, arr2)
}

func merge(arr1, arr2 []int) (ret []int) {
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] > arr2[0] {
			ret = append(ret, arr2[0])
			arr2 = arr2[1:]
		} else {
			ret = append(ret, arr1[0])
			arr1 = arr1[1:]
		}
	}
	if len(arr1) == 0 {
		ret = append(ret, arr2...)
	} else {
		ret = append(ret, arr1...)
	}
	return
}
