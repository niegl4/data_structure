package array

import "sort"

/*
六十一
判断五张扑克牌是否是顺子。A为1，JQK为11，12，13，大小王为0。

1.排序
2.计算0的个数
3.计算gap的和，并校验是否有相等的非0元素
4.比较0的和与gap的和
*/
func isContinuous(arr []int) bool {
	length := len(arr)
	if length < 1 {
		return false
	}
	if length == 1 {
		return true
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	numOfZero := 0
	numOfGap := 0
	for i := 0; i < length; i++ {
		if arr[i] == 0 {
			numOfZero++
		}
	}

	for i := 0; i < length-1; i++ {
		if arr[i] == arr[i+1] && arr[i] != 0 {
			return false
		}
		sub := arr[i+1] - arr[i]
		if sub > 1 {
			numOfGap += sub - 1
		}
	}

	if numOfGap > numOfZero {
		return false
	} else {
		return true
	}
}
