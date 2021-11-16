package _4_merge

//先切分，再排序
//最后从最小的单位逐级返回，每次返回都是有序的
func MergeSort(a []int, n int) {
	if n <= 1 {
		return
	}
	mergeSort(a, 0, n-1)
}

func mergeSort(a []int, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	mergeSort(a, p, q)
	mergeSort(a, q+1, r)
	merge(a, p, q, r)
}

func merge(a []int, p, q, r int) {
	i := p //左区间索引
	j := q + 1 //右区间索引
	tmp := make([]int, 0, r-p+1) //临时数组，就是因为它，导致空间复杂度O(n)
	//合并两个有序数组
	for i <= q && j <= r {
		if a[i] > a[j] {
			tmp = append(tmp, a[j])
			j++
		} else {
			tmp = append(tmp, a[i])
			i++
		}
	}
	//收集剩余元素
	if i <= q {
		for ; i <= q; i++ {
			tmp = append(tmp, a[i])
		}
	} else {
		for ; j <= r; j++ {
			tmp = append(tmp, a[j])
		}
	}
	//临时数组拷贝进原数组，注意索引偏移
	for k := 0; k <= r-p; k++ {
		a[p+k] = tmp[k]
	}
}