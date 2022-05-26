package bit_calculate

/*
十五-1
*
给定一个整数，输出该数的二进制表示中1的个数。

原数可能是正整数，也可能是负整数，有符号数的右移会有两种情况，不好处理。
问题转换为：游标为1，每次左移1，再与原数位与运算。左移的情况简单，所以简化了问题。
*/
func numberOf1Version1(n int) int {
	count := 0
	flag := 1
	for flag != 0 {
		if n&flag != 0 {
			count++
		}
		flag <<= 1
	}
	return count
}

//【把一个整数减去1，再和原整数做与运算，会把该整数的最右边的1变成0】
func numberOf1Version2(n int) int {
	count := 0
	for n != 0 {
		n = (n - 1) & n
		count++
	}
	return count
}

/*
十五-2
用1条语句判断一个整数是不是2的自然数次方。
*/
func judgeNumber(n int) bool {
	return (n-1)&n == 0
}

/*
十五-3
输入两个整数m，n。需要改变m的二进制表示中的多少位，才能得到n
*/
func changeMtoN(m, n int) int {
	diff := m ^ n
	return numberOf1Version2(diff)
}
