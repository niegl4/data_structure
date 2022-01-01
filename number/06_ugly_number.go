package number

/*
四十九
只包含因子2,3,5的数称为丑数。求从小到大第1500个丑数。第一个丑数为1。
 */
func getUglyNumberV1(index int) int {
	if index <= 0 {
		return 0
	}

	number := 0
	uglyCount := 0
	for uglyCount < index {
		number++
		if isUgly(number) {
			uglyCount++
		}
	}
	return number
}

func isUgly(number int) bool {
	for number % 2 == 0 {
		number /= 2
	}
	for number % 3 == 0 {
		number /= 3
	}
	for number % 5 == 0 {
		number /= 5
	}
	return number == 1
}

func getUglyNumberV2(index int) int {
	if index <= 0 {
		return 0
	}

	uglyNumSet := make([]int, index)
	uglyNumSet[0] = 1

	uglyCount := 1
	multiply2 := 0
	multiply3 := 0
	multiply5 := 0
	for uglyCount < index {
		uglyNumSet[uglyCount] = minOf3Nums(uglyNumSet[multiply2]*2, uglyNumSet[multiply3]*3, uglyNumSet[multiply5]*5)
		for uglyNumSet[multiply2] * 2 <= uglyNumSet[uglyCount] {
			multiply2++
		}
		for uglyNumSet[multiply3] * 3 <= uglyNumSet[uglyCount] {
			multiply3++
		}
		for uglyNumSet[multiply5] * 5 <= uglyNumSet[uglyCount] {
			multiply5++
		}
		uglyCount++
	}
	return uglyNumSet[index-1]
}

func minOf3Nums(num1, num2, num3 int) int {
	min := num1
	if num2 < num1 {
		min = num2
	}
	if num3 < min {
		min = num3
	}
	return min
}
