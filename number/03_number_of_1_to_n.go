package number

import (
	"fmt"
	"strconv"
)

/*
四十三
1~n个整数的十进制表示中，1出现的次数。
如1~12中包含1的数字1，10，11，12，1出现了共五次。
*****
*/
func numberOf1ToN(n int) int {
	if n <= 0 {
		return 0
	}
	str := fmt.Sprintf("%d", n)
	var arr []int
	for _, b := range []byte(str) {
		num := int(b - '0')
		arr = append(arr, num)
	}
	return numberOf1(arr)
}

func numberOf1(number []int) int {
	if len(number) < 1 {
		return 0
	}
	first := number[0] - 0

	if len(number) == 1 && first == 0 {
		return 0
	}
	if len(number) == 1 && first > 0 {
		return 1
	}

	numFirstDigit := 0
	if first > 1 {
		numFirstDigit = powerBase10(len(number) - 1)
	} else {
		str := ""
		for i := 1; i < len(number); i++ {
			str += fmt.Sprintf("%d", number[i])
		}
		numFirstDigit, _ = strconv.Atoi(str)
		numFirstDigit++
	}

	numOtherDigits := first * (len(number) - 1) * powerBase10(len(number)-2)
	numRecursive := numberOf1(number[1:])
	return numFirstDigit + numOtherDigits + numRecursive
}

func powerBase10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}
