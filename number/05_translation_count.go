package number

import (
	"fmt"
	"strconv"
)

/*
四十六
给定整数，计算翻译的方式有多少种。0翻译为a,25翻译为z。
*

把整数转换成字符串，把字符串转换成[]int32
*/

func translationCount(number int) int {
	if number < 0 {
		return 0
	}
	numberStr := fmt.Sprintf("%d", number)
	length := len(numberStr)
	numberSli := make([]int32, length)
	count := make([]int, length)
	for _, n := range numberStr {
		a := n - '0'
		numberSli = append(numberSli, a)
	}

	for i := length - 1; i >= 0; i-- {
		if i == length-1 {
			count[i] = 1
		} else {
			count[i] = count[i+1]
		}

		if i <= length-3 {
			num, _ := strconv.Atoi(numberStr[i : i+2])
			if num >= 0 && num <= 25 {
				count[i] += count[i+2]
			}
		}
	}

	return count[0]
}
