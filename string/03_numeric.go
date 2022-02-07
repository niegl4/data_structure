package string

/*
二十
实现一个函数判断字符串是否表示数值。

规则如下：
A[.[B]][e|EC]
.B[e|EC]
A:整数部分，以+或-开头的0~9的数位串；可以有符号，也可以没有符号。
B:小数部分，0~9的数位串；不能有符号。
C:指数部分，以+或-开头的0~9的数位串；可以有符号，也可以没有符号。
*/
func isNumeric(str string) bool {
	if str == "" {
		return false
	}
	strSli := make([]rune, 0)
	for _, char := range str {
		strSli = append(strSli, char)
	}

	var (
		numeric1, numeric2, numeric3 bool
		idx                          int
		idxMax                       = len(strSli)
	)

	numeric1, idx = scanInt(strSli, 0)
	if idx >= idxMax {
		return numeric1
	}

	if strSli[idx] == '.' {
		idx++

		if idx >= idxMax {
			return numeric1
		}
		numeric2, idx = scanUint(strSli, idx)
		if !(numeric1 || numeric2) {
			return false
		} else {
			if idx >= idxMax {
				return true
			} else {
				if strSli[idx] == 'e' || strSli[idx] == 'E' {
					idx++

					if idx >= idxMax {
						return false
					}
					numeric3, idx = scanInt(strSli, idx)
				}
			}
		}
	}
	if strSli[idx] == 'e' || strSli[idx] == 'E' {
		idx++
		if idx >= idxMax {
			numeric3 = false
		} else {
			numeric3, idx = scanInt(strSli, idx)
		}
	}
	return (numeric1 || numeric2) && numeric3 && idx >= idxMax
}

func scanUint(strSli []rune, idx int) (bool, int) {
	target := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	n := len(strSli)
	index := idx
	for ; index < n; index++ {
		flag := false
		for _, char := range target {
			if strSli[index] == char {
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return index > idx, index
}

func scanInt(strSli []rune, idx int) (bool, int) {
	if strSli[idx] == '+' || strSli[idx] == '-' {
		idx++
	}
	return scanUint(strSli, idx)
}
