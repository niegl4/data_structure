package array

/*
六十三
股票的价格按时间顺序存储在数组中。求最大利润。
*/

func maxDiff(arr []int) int {
	length := len(arr)
	if length < 2 {
		return 0
	}

	min := arr[0]
	maxDiff := arr[1] - min

	for i := 2; i < length; i++ {
		if arr[i-1] < min {
			min = arr[i-1]
		}

		curDiff := arr[i] - min
		if curDiff > maxDiff {
			maxDiff = curDiff
		}
	}
	return maxDiff
}
