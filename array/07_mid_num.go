package array

import "errors"

/*
三十九
数组中出现次数超过一半的数字。
*/

//超过一半的数字，也即中位数
func midNumByPartition(arr []int) (int, error) {
	n := len(arr)
	if n <= 0 {
		return 0, errors.New("no mid num")
	}
	if n == 1 {
		return arr[0], nil
	}

	//初始分区为整个数组
	midNumIdx := n >> 1
	low := 0
	up := n - 1
	resultIdx := 0
	for {
		pivot := arr[up]
		//注意：i，j不同步。循环结束，i就是arr[up]的位置
		i := low
		for j := low; j < up; j++ {
			if arr[j] > pivot {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
		}
		arr[i], arr[up] = arr[up], arr[i]

		if i+1 == midNumIdx { //循环退出的情况
			resultIdx = i
			break
		} else if midNumIdx <= i { //k出现在左半边，减小上限，继续搜索
			up = i - 1
		} else { //k出现在右半边，增大下限，继续搜索
			low = i + 1
		}
	}
	if !checkMoreThanHalf(arr, arr[resultIdx]) {
		return 0, errors.New("mid num not exist")
	}
	return arr[midNumIdx], nil
}

func checkMoreThanHalf(arr []int, num int) bool {
	times := 0
	for _, item := range arr {
		if item == num {
			times++
		}
	}
	return times*2 > len(arr)
}

//***计数实现。
func midNumBySkill(arr []int) (int, error) {
	n := len(arr)
	if n <= 0 {
		return 0, errors.New("no mid num")
	}
	if n == 1 {
		return arr[0], nil
	}

	result := arr[0]
	times := 1
	for i := 1; i < n; i++ {
		if times == 0 {
			result = arr[i]
			times = 1
		} else if arr[i] == result {
			times++
		} else {
			times--
		}
	}

	if !checkMoreThanHalf(arr, result) {
		return 0, errors.New("mid num not exist")
	}
	return result, nil
}
