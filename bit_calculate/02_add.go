package bit_calculate

/*
六十五-1
不用+，-，*，/，做加法。
*/
func add(num1, num2 int) int {
	sum := num1 ^ num2
	carry := (num1 & num2) << 1
	num1 = sum
	num2 = carry
	for num2 != 0 {
		sum = num1 ^ num2
		carry = (num1 & num2) << 1
		num1 = sum
		num2 = carry
	}
	return num1
}

/*
六十五-2
不使用新变量，交换两个变量的值。
*/
func swapV1(num1, num2 int) (int, int) {
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	return num1, num2
}

func swapV2(num1, num2 int) (int, int) {
	num1 = num1 ^ num2
	num2 = num1 ^ num2
	num1 = num1 ^ num2
	return num1, num2
}
