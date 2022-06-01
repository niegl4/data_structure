package bit_calculate

/*
六十五-1
不用+，-，*，/，做加法。

*
1.不考虑进位对每一位相加。刚好是异或的规则。
2.计算进位。刚好是与的规则，再加左移一位。
3.直到不产生进位为止。
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

*
尤其是通过异或的方式，就是利用"一个数异或它本身等于0，一个数异或0等于它本身"。
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
