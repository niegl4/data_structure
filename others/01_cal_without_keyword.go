package others

/*
六十四
不用乘除法，不用语法关键字，计算：1+2+...+n

函数指针求解
*/
type sumWithOutKeyword func(int) int

func terminator(n int) int {
	return 0
}

func sum(n int) int {
	fMap := map[bool]sumWithOutKeyword{
		false: sum,
		true:  terminator,
	}
	return n + fMap[n == 0](n-1)
}
