package array

/*
六十六
给定数组，构造数组，不使用乘法。
*/
func multiply(arr []int) []int {
	length := len(arr)
	if length < 1 {
		return nil
	}

	tmp1 := make([]int, length)
	tmp2 := make([]int, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			tmp1[i] = 1
			tmp2[length-1-i] = 1
		} else {
			tmp1[i] = tmp1[i-1] * arr[i-1]
			tmp2[length-1-i] = tmp2[length-i] * arr[length-i]
		}
	}

	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = tmp1[i] * tmp2[i]
	}
	return res
}
