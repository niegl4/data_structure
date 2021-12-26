package number

/*
四十四
数字序列中的某一位的数字。求第n位的数字，n从0开始计数。
0~15 =》"0123456789101112131415"
*/

func digitAtIndex(index int) int {
	if index < 0 {
		return -1
	}
	digits := 1
	for {
		numbers := countOfInt(digits)
		if index < numbers*digits {
			return digitAtIndexCore(index, digits)
		}
		index -= digits * numbers
		digits++
	}
	return -1
}

func countOfInt(digits int) int {
	if digits == 1 {
		return 10
	}
	count := powerBase10(digits - 1)
	return 9 * count
}

func digitAtIndexCore(idx, digits int) int {
	number := beginNum(digits) + idx/digits
	indexFromRigth := digits - idx%digits
	for i := 1; i < indexFromRigth; i++ {
		number /= 10
	}
	return number % 10
}

func beginNum(digits int) int {
	if digits == 1 {
		return 0
	}
	res := 1
	for i := 1; i < digits; i++ {
		res *= 10
	}
	return res
}
