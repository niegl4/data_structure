package number

/*
四十四
*
数字序列中的某一位的数字。求第n位的数字，n从0开始计数。
0~15 =》"0123456789101112131415"
*/

//求digits位的数字总共有几个，eg：1位数，10个；2位数，90个；3位数，900个
func countOfInt(digits int) int {
	if digits == 1 {
		return 10
	}
	count := powerBase10(digits - 1)
	return 9 * count
}

func digitAtIndex(index int) int {
	if index < 0 {
		return -1
	}
	digits := 1
	for {
		numbers := countOfInt(digits)
		if index < numbers*digits {
			return digitAtIndexCoreV2(index, digits)
		}
		index -= digits * numbers
		digits++
	}
	return -1
}

//eg: digitAtIndexCoreV2(811, 3)
func digitAtIndexCoreV2(idx, digits int) int {
	count := idx / digits                 // 811 = [270] * 3 + 1
	loopCnt := idx - count*digits         // 811 = 270 * 3 + [1]
	targetNum := beginNum(digits) + count // 100+270，其实是第270个数的下一个数
	return getNum(targetNum, digits, loopCnt)
}

//digits位的数字，第一个是多少，eg：1位数，0；2位数，10；3位数，100
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

//已知一个数字，求它第几位上的数字
func getNum(targetNum, digits, loopCnt int) (num int) {
	//循环次数由bit确定
	for i := 0; i < loopCnt+1; i++ {
		supNum := powerBase10(digits - 1)
		digits--
		num = targetNum / supNum
		targetNum %= supNum
	}
	return num
}

////书上的写法，难以看懂
//func digitAtIndexCore(idx, digits int) int {
//	number := beginNum(digits) + idx/digits
//	indexFromRight := digits - idx%digits
//	for i := 1; i < indexFromRight; i++ {
//		number /= 10
//	}
//	return number % 10
//}
