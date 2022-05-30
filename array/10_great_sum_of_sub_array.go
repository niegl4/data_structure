package array

/*
四十二
*
连续子数组的最大和。

1.当前和为负数，抛弃当前和，从arr[i]开始求和
2.当前和为正数，不管arr[i]正或负，都要加上（因为即时为负，后续可能会有更大的正数）
3.函数内设置一个临时max，当前和与它比较并刷新

暴力求解，O(n^2)个子数组，时间复杂度为O(n^2)
设置临时max，O(n)
*/
func findGreatSumOfSubArray(arr []int) int {
	length := len(arr)
	if length <= 0 {
		return 0
	}

	var (
		curSum      = 0
		greatestSum = -0x8000000000000000 //int64的最小值
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
