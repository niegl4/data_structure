package array

import "errors"

/*
五十六-1
数组中只出现一次的两个数字。
一个整形数组中，除了两个数字外，其他数字都出现了两次。
*/
func twoNumAppearOnce(arr []int) (int, int, error) {
	length := len(arr)
	if length < 2 {
		return 0, 0, errors.New("input array invalid")
	}
	num := 0
	for _, item := range arr {
		num ^= item
	}

	bitIndex := findFirstBitIs1FromRight(num)
	num1 := 0
	num2 := 0
	for _, item := range arr {
		if isBit1(item, bitIndex) {
			num1 ^= item
		} else {
			num2 ^= item
		}
	}
	return num1, num2, nil
}

func findFirstBitIs1FromRight(num int) int {
	firstBit := 0
	for num&0b1 == 0 && firstBit < 64 {
		num >>= 1
		firstBit++
	}
	return firstBit
}

func isBit1(num int, bitIndex int) bool {
	num >>= bitIndex
	return num&0b1 == 1
}

/*
五十六-2
数组中只出现一次的一个数字。
一个整形数组中，除了一个数字外，其他数字都出现了三次。
*/
func oneNumAppearOnce(arr []int) int {
	length := len(arr)

	bitSum := make([]int, 64)
	for i := 0; i < length; i++ {

		bitMask := 0b1
		for j := 63; j >= 0; j-- {
			if arr[i]&bitMask != 0 {
				bitSum[j] += 1
			}
			bitMask <<= 1
		}
	}

	result := 0
	for i := 0; i < 64; i++ {
		result <<= 1
		result += bitSum[i] % 3
	}
	return result
}
