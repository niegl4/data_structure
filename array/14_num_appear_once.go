package array

import "errors"

/*
五十六-1
数组中只出现一次的两个数字。
一个整形数组中，除了两个数字外，其他数字都出现了两次。
*
1.任意数字，异或它本身，等于0.
2.把数组分成两组，每组各包含一个只出现一次的数字。异或就能得到该数字。
3.怎么确定分组依据？原数组异或，右边开始第一个bit位为1，说明那两个数字在该bit位不等。以此为依据。
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
	mask := 0b1
	if bitIndex > 0 {
		mask <<= bitIndex
	}
	num1 := 0
	num2 := 0
	for _, item := range arr {
		if isBit1(item, mask) {
			num1 ^= item
		} else {
			num2 ^= item
		}
	}
	return num1, num2, nil
}

func findFirstBitIs1FromRight(num int) int {
	firstBit := 0
	mask := 0b1
	for num&mask == 0 && firstBit < 64 {
		mask <<= 1
		firstBit++
	}
	return firstBit
}

func isBit1(num int, mask int) bool {
	return num&mask > 0
}

/*
五十六-2
数组中只出现一次的一个数字。
一个整形数组中，除了一个数字外，其他数字都出现了三次。
*
bit位进行累计统计，对3求余数，余数就代表了该数字在该bit位是否为1.（余数其实要么为0，要么为1）
*/
func oneNumAppearOnce(arr []int) int {
	length := len(arr)

	bitSum := make([]int, 64) //索引0，对应数字的高位；索引63，对应数字的低位。
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
		if i == 0 {
			result += bitSum[i] % 3
		} else {
			result <<= 1 //共64位，只需要移位63次
			result += bitSum[i] % 3
		}
	}
	return result
}
