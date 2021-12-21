package array

/*
二十一
输入一个整数数组，实现函数：调整数字顺序，奇数位于前半部分，偶数位于后半部分。
 */
func reorder(arr []int, f func (num int) bool) {
	i := 0
	n := len(arr)
	for j := 0; j < n; j++ {
		if f(arr[j]) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
}

func isOdd(num int) bool {
	return num & 0b1 == 1
}

func isNegative(num int) bool {
	return num < 0
}

func isDivisibleBy3(num int) bool {
	return num % 3 == 0
}
