package search

import "errors"

/*
十一-1
递增数组的旋转数组，找到最小值。

二分查找的变种，时间复杂度O(logn)。
 */
func searchMinInRotateSortArr1(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("invalid input")
	}
	if len(arr) == 1 {
		return arr[0], nil
	}

	start := 0
	end := len(arr) - 1
	mid := start
	for arr[start] >= arr[end] {//没有旋转，arr[0]就是min
		if end - start == 1 {
			mid = end
			break
		}

		mid = (end - start) >> 2 + start
		if arr[start] == arr[end] && arr[start] == arr[mid] { //首，尾，中全都相等，无法缩小区间，只能顺序遍历
			min := arr[start]
			for i := start+1; i <= end; i++ {
				if arr[i] < min {
					min = arr[i]
				}
			}
			return min, nil
		}

		if arr[mid] >= arr[start] {
			start = mid
		} else {
			end = mid
		}
	}
	return arr[mid], nil
}

//通过索引缩小查找边界，可读性更好
func searchMinInRotateSortArr2(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("invalid input")
	}
	if len(arr) == 1 {
		return arr[0], nil
	}

	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (end - start) >> 2 + start

		if arr[start] < arr[end] { //全局递增数据，相当于没有任何旋转
			return arr[start], nil
		} else if arr[start] > arr[end] {
			if arr[mid] > arr[start] {
				start = mid + 1
			} else if arr[mid] == arr[start] {
				start = mid + 1
			} else {
				if arr[mid-1] > arr[mid] {
					return arr[mid], nil
				} else {
					end = mid - 1
				}
			}
		} else {
			if arr[mid] > arr[start] {
				start = mid + 1
			} else if arr[mid] == arr[start] { //首，中，尾三个相等，无法缩小边界，遍历查找
				min := arr[start]
				for i := start+1; i <= end; i++ {
					if arr[i] < min {
						min = arr[i]
					}
				}
				return min, nil
			} else {
				if arr[mid-1] > arr[mid] {
					return arr[mid], nil
				} else {
					end = mid - 1
				}
			}
		}
	}

	return 0, errors.New("exception condition occur")
}
