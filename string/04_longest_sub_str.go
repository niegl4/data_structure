package string

import "errors"

/*
四十八
从字符串中找出一个最长的不包含重复字符的子字符串，计算该子字符串的长度。
字符串中只包含“a”~"z"的字符。
*/
func lengthOfLongestSubStr(str string) (int, error) {
	length := len(str)
	if length <= 1 {
		return length, nil
	}

	position := make([]int, 26)
	for idx := range position {
		position[idx] = -1
	}

	var (
		curLen = 0
		maxLen = 0
	)
	for idx, char := range []byte(str) {
		charNum := char - 'a'
		if charNum < 0 || charNum > 25 {
			return 0, errors.New("str invalid")
		}

		if position[charNum] < 0 || (position[charNum] >= 0 && idx-position[charNum] > curLen) {
			curLen++
		} else {
			curLen = idx - position[charNum]
		}

		position[charNum] = idx
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen, nil
}

/*
数组只包含0,1，找出最长的连续1
 */
func check1(arr []int, idx int) bool {
	return arr[idx] == 1
}

func check0(arr []int, idx int) bool {
	return arr[idx] == 0
}

func longestIndex(arr []int, f func ([]int, int) bool) (idx, maxLength int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		if f(arr, i) && i >= idx + maxLength { //优化，当前i没有逃出idx+maxLength，则不用后续判断
			j := i
			for ; j < length && f(arr, j); j++ {
			}
			if j - i > maxLength {
				maxLength = j - i
				idx = i
			}
		}
	}
	if maxLength == 0 {
		idx = -1
	}
	return idx, maxLength
}
