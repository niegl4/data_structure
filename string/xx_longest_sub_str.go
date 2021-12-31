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

		if position[charNum] < 0 || (position[charNum] >= 0 && idx - position[charNum] > curLen) {
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
