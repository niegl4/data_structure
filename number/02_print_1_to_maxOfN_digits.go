package number

import (
	"fmt"
)

/*
十七
*
输入数字n，按顺序打印从1到最大的n位十进制数。
题中没有显式提醒"不用考虑大数"，那就暗示这是一个大数问题，用string表示大数。

"全排列"解法更适合面试写。
*/

func print1ToMaxOfDigits(n int) {
	if n <= 0 {
		return
	}

	numberStr := make([]byte, n)
	for i := 0; i < n; i++ {
		numberStr[i] = 48
	}
	for !increment(numberStr) {
		printNumber(numberStr)
	}
}

//加1时，第一个字符产生了进位，则表明当前已经是最大的n位数，return true
func increment(numberStr []byte) bool {
	var (
		isOverflow = false
		takeOver   = uint8(0)
		maxLen     = len(numberStr)
	)
	for i := maxLen - 1; i >= 0; i-- {
		bitSum := numberStr[i] - '0' + takeOver
		if i == maxLen-1 {
			bitSum++
		}

		if bitSum >= 10 {
			if i == 0 {
				isOverflow = true
			} else {
				bitSum -= uint8(10)
				takeOver = uint8(1)
				numberStr[i] = '0' + bitSum
			}
		} else {
			numberStr[i] = '0' + bitSum
			break
		}
	}
	return isOverflow
}

func printNumber(numberStr []byte) {
	maxLen := len(numberStr)
	for i := 0; i < maxLen; i++ {
		if i == maxLen-1 {
			fmt.Println(string(numberStr[maxLen-1]))
			return
		} else {
			if numberStr[i] == '0' {
				continue
			} else {
				fmt.Println(string(numberStr[i:]))
				return
			}
		}
	}
	return
}

//每个数位0~9的全排列
func print1ToMaxOfDigitsV2(n int) {
	if n <= 0 {
		return
	}

	numberStr := make([]byte, n)
	for i := 0; i < n; i++ {
		numberStr[i] = 48
	}

	for i := 0; i < 10; i++ {
		numberStr[0] = 48 + uint8(i) //这里的下标0，位于切片的最左端，是大数的最高位。
		print1ToMaxOfDigitsV2Core(numberStr, 0)
	}
}

func print1ToMaxOfDigitsV2Core(numberStr []byte, index int) {
	if index == len(numberStr)-1 {
		printNumber(numberStr)
		return
	}

	for i := 0; i < 10; i++ {
		numberStr[index+1] = 48 + uint8(i)
		print1ToMaxOfDigitsV2Core(numberStr, index+1)
	}
}
