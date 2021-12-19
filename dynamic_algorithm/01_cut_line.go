package dynamic_algorithm

import "errors"

/*
十四-1
n长的绳子，剪成m段（m，n都是整数，n>1，m>1）。剪之后的绳子乘积最大是多少。
绳子长度至少为2，至少剪1刀。

动态规划解，时间复杂度O(n^2).
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

	product := make([]int, lineLen+1)
	product[0] = 0
	product[1] = 1
	product[2] = 2
	product[3] = 3

	for i := 4; i <= lineLen; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			tmp := product[j] * product[i-j]
			if tmp > max {
				max = tmp
			}
		}
		product[i] = max
	}
	return product[lineLen], nil
}
