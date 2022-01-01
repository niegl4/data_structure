package array

/*
五十一
统计数组中的逆序对。
*/

func inversePairs(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	return inversePairsCore(arr, 0, len(arr)-1)
}

func inversePairsCore(arr []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := (end-start)>>2 + start
	left := inversePairsCore(arr, start, mid)
	right := inversePairsCore(arr, mid+1, end)
	count := merge(arr, start, mid, end)
	return left + right + count
}

func merge(arr []int, start, mid, end int) (count int) {
	i := mid                        //左区间索引终点
	j := end                        //右区间索引终点
	tmp := make([]int, end-start+1) //临时数组，就是因为它，导致空间复杂度O(n)
	tmpIdx := end - start
	//合并两个有序数组
	for idx := end; idx >= start && i >= start && j >= mid+1; idx-- {
		if arr[i] > arr[j] {
			count += j - (mid + 1) + 1

			tmp[tmpIdx] = arr[i]
			i--
		} else {
			tmp[tmpIdx] = arr[j]
			j--
		}
		tmpIdx--
	}
	//收集剩余元素
	for ; i >= start; i-- {
		tmp[tmpIdx] = arr[i]
		tmpIdx--
	}
	for ; j >= mid+1; j-- {
		tmp[tmpIdx] = arr[j]
		tmpIdx--
	}

	//临时数组拷贝进原数组，注意索引偏移
	for k := 0; k <= end-start; k++ {
		arr[start+k] = tmp[k]
	}
	return count
}
