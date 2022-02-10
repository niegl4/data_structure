package search

import "errors"

/*
五十三-1
统计一个数字在排序数组中出现的次数。
*/
func countInSortArray(arr []int, k int) int {
	if len(arr) < 1 {
		return 0
	}
	start := BSearchExtend1(arr, k)
	if start == -1 {
		return 0
	}
	end := BSearchExtend2(arr, k)
	return end - start + 1
}

/*
五十四-2
n-1长度的排序数组，每个数字都是唯一的，取值范围是0~n-1。求唯一一个在取值范围但不在数组中的数。
*/
func missNumInSortArray(arr []int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := (high-low)>>1 + low
		if arr[mid] == mid {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] == mid-1 {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

/*
五十五-3
一个单调递增的数组里，每个元素都是整数且唯一。求任一数值等于下标的元素。
*/
func numberEqualIndex(arr []int) (int, error) {
	n := len(arr)
	if n < 1 {
		return 0, errors.New("input array invalid")
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := (high-low)>>1 + low
		if arr[mid] == mid {
			return mid, nil
		} else if arr[mid] < mid {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0, errors.New("input array invalid")
}
