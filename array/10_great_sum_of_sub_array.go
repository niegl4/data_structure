package array

/*
四十二
连续子数组的最大和。
*/
func findGreatSumOfSubArray(arr []int) int {
	length := len(arr)
	if length <= 0 {
		return 0
	}

	var (
		curSum      = 0
		greatestSum = -0x8000000000000000
	)
	for i := 0; i < length; i++ {
		if curSum <= 0 {
			curSum = arr[i]
		} else {
			curSum += arr[i]
		}
		if curSum > greatestSum {
			greatestSum = curSum
		}
	}
	return greatestSum
}
