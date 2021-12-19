package greedy_algorithm

import "errors"

/*
十四-2
n长的绳子，剪成m段（m，n都是整数，n>1，m>1）。剪之后的绳子乘积最大是多少。
绳子长度至少为2，至少剪1刀。

贪婪解，时间复杂度O(n)。但需要数学证明。
*/
func maxProductAfterCut(lineLen int) (int, error) {
	if lineLen < 2 {
		return 0, errors.New("input invalid")
	}
	if lineLen == 2 {
		return 1, nil
	}
	if lineLen == 3 {
		return 2, nil
	}

	timesOf3 := lineLen / 3
	if lineLen % timesOf3*3 == 1 {
		timesOf3--
	}

	timesOf2 := (lineLen - timesOf3*3) / 2

	max := 1
	for i := 0; i < timesOf3; i++ {
		max *= 3
	}
	for i := 0; i < timesOf2; i++ {
		max *= 2
	}

	return max, nil
}
