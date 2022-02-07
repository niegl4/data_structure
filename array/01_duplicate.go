package array

import "errors"

/*
三
n长数组里所有数字都在0~n-1，可以改变数组排序，找出任一重复数字

每个数字最多交换2次，O(n)
*/
func duplicate1(arr []int) (dupNum int, err error) {
	length := len(arr)
	if length <= 1 {
		return 0, errors.New("duplicate number not exits")
	}

	for _, num := range arr {
		if num < 0 || num > length-1 {
			return 0, errors.New("array invalid")
		}
	}

	for i, num := range arr {
		for i != num { //只要索引不等于值，就持续交换
			if arr[num] == num { //num位，索引等于值，就说明该数字重复
				return num, nil
			}
			arr[i], arr[num] = arr[num], arr[i]
		}
	}
	return 0, errors.New("duplicate number not exits")
}

/*
n+1长数组里所有数字都在1~n，不能改变数组排序，找出任一重复数字.
这里二分查找比较巧妙：数组本身无序，但是数据区间是有序的，即1~n。

二分查找O(logn)，每次查找中都需要遍历数组统计数量O(n)，总共O(nlogn)
*/
func duplicate2(arr []int) (dupNum int, err error) {
	length := len(arr)
	if length <= 1 {
		return 0, errors.New("duplicate number not exits")
	}

	low := 1
	high := length - 1
	for _, num := range arr {
		if num < low || num > high {
			return 0, errors.New("array invalid")
		}
	}

	calCount := func(start, end int) (countNum int) {
		for _, num := range arr {
			if num >= start && num <= end {
				countNum++
			}
		}
		return countNum
	}

	for low <= high {
		mid := (high-low)>>2 + low
		count := calCount(low, mid)

		if low == high {
			if count > 1 {
				return low, nil
			} else {
				break
			}
		}

		numCount := mid - low + 1
		if count == numCount {
			low = mid + 1
		} else if count > numCount {
			high = mid //这里边界条件的修改，为啥不是high = mid - 1？
		} else {
			low = mid + 1
		}
	}
	return 0, errors.New("duplicate number not exits")
}
