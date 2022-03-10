package _4_merge

//先切分，再排序
//最后从最小的单位逐级返回，每次返回都是有序的
func MergeSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	mergeSort(arr, 0, length-1)
}

func mergeSort(arr []int, p, r int) {
	if p >= r {
		return
	}
	mid := (p + r) / 2
	mergeSort(arr, p, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, p, mid, r)
}

func merge(arr []int, p, mid, r int) {
	i := p                       //左区间索引
	j := mid + 1                   //右区间索引
	tmp := make([]int, 0, r-p+1) //临时数组，就是因为它，导致空间复杂度O(n)
	//合并两个有序数组
	for i <= mid && j <= r {
		if arr[i] > arr[j] {
			tmp = append(tmp, arr[j])
			j++
		} else {
			tmp = append(tmp, arr[i])
			i++
		}
	}
	//收集剩余元素.上方循环退出的条件是：i，j其中一个大于了上限
	if i <= mid {
		for ; i <= mid; i++ {
			tmp = append(tmp, arr[i])
		}
	} else if j <= r {
		for ; j <= r; j++ {
			tmp = append(tmp, arr[j])
		}
	}
	//临时数组拷贝进原数组，注意索引偏移
	for k := 0; k < r-p+1; k++ {
		arr[p+k] = tmp[k]
	}
}
