package bit_calculate

/*
六十五
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
