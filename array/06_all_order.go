package array

import (
	"strconv"
	"strings"
)

/*
三十八-1
输出字符数组的全排列。

（number的第十七，也有排列问题。）
*/
func allOrder(arr []string) (res []string) {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	return allOrderCore(arr, 0)
}

//分为两个子问题：idx索引位的确定（它以及它之后的每个元素轮流交换）+ 数组排除idx之后的全排列
//此处是：for i := idx; i < length; i++ ，idx为形参，每次调用递增idx
//也可以换成：for i := 0; i < length; i++ ，length为形参，每次调用缩小length
func allOrderCore(arr []string, idx int) (res []string) {
	length := len(arr)
	if idx == length-1 {
		return []string{strings.Join(arr, "|")}
	}
	for i := idx; i < length; i++ {
		arr[idx], arr[i] = arr[i], arr[idx]
		res = append(res, allOrderCore(arr, idx+1)...)
		arr[idx], arr[i] = arr[i], arr[idx]
	}
	return res
}

func allOrderCoreV2(arr []string, length int) (allOrderSet []string) {
	if length == 1 {
		return []string{strings.Join(arr, "|")}
	}

	for i := 0; i < length; i++ {
		arr[i], arr[length-1] = arr[length-1], arr[i]
		allOrderSet = append(allOrderSet, allOrderCoreV2(arr, length-1)...)
		arr[i], arr[length-1] = arr[length-1], arr[i]
	}
	return allOrderSet
}

/*
三十八-2
输出字符数组的全组合。
[a, b, c] => [a, b, c, a|b, a|c, b|c, a|b|c]
*/
func allCombination(arr []string) (res []string) {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	//长度从1到len，逐个计算
	for targetLen := 1; targetLen <= length; targetLen++ {
		res = append(res, allCombinationCore(arr, 0, targetLen)...)
	}
	return res
}

//以[a, b, c, d]求长度为2的组合为例，推边界条件。
func allCombinationCore(arr []string, start, targetLen int) (res []string) {
	//优化的写法，但是这样，就不没有考虑边界条件了。
	if targetLen == 1 {
		res = append(res, arr[start:]...)
		return
	}
	if targetLen == len(arr[start:]) {
		return []string{strings.Join(arr[start:], "|")}
	}

	part1 := allCombinationCore(arr, start+1, targetLen-1)
	for _, str := range part1 {
		res = append(res, arr[start]+"|"+str)
	}

	res = append(res, allCombinationCore(arr, start+1, targetLen)...)
	return res
}

/*
三十八-3
*/
func numsOfQueens(nxn int) (res [][]int) {
	if nxn <= 0 {
		return nil
	}
	if nxn == 1 {
		return [][]int{
			{0},
		}
	}
	if nxn == 2 {
		return nil
	}

	//已有的全排列，是对string对象。为了能调用它，临时生成"0"，"1"这样的string。
	matrix := make([]string, nxn)
	for i := 0; i < nxn; i++ {
		matrix[i] = strconv.Itoa(i)
	}

	allOrders := allOrder(matrix)
	length := len(allOrders)
check:
	for i := 0; i < length; i++ {
		tmp := strings.Split(allOrders[i], "|")
		tmpNum := make([]int, nxn)
		for j := 0; j < nxn; j++ {
			tmpNum[j], _ = strconv.Atoi(tmp[j])
		}

		for m := 0; m < nxn; m++ {
			for n := m + 1; n < nxn; n++ {
				sub := m - n
				if tmpNum[m]-tmpNum[n] == sub || tmpNum[n]-tmpNum[m] == sub {
					continue check
				}
			}
		}
		res = append(res, tmpNum)
	}
	return res
}
