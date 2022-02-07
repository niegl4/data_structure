package recursion

import "errors"

/*
十-1
求斐波那契数列的第n项

递归实现的时间复杂度为2^n。
循环实现的时间复杂度为O(n)。

跳台阶的问题，是一个类斐波那契数列，它和斐波那契数列不是完全相同。
小矩形覆盖大矩形的问题，也是一个类斐波那契数列。
*/
func fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("input invalid")
	}
	fiArr := []int{0, 1}
	if n <= 1 {
		return fiArr[n], nil
	}

	var (
		num1 = 0
		num2 = 1
		fiN  = 0
	)

	for i := 2; i <= n; i++ {
		fiN = num1 + num2
		if fiN < 0 { //小于0，说明整数溢出
			return 0, errors.New("int overflow")
		} else {
			num1 = num2
			num2 = fiN
		}
	}
	return fiN, nil
}
