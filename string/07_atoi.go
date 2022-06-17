package string

import "errors"

/*
六十七
string转int。

*
1.str invalid：1.为空；2.正负号开头，正数的话也可以省略正号；3.str中存在非数字的字符
2.str是一个大数，对int溢出。可以是正溢出，也可以是负溢出。
*/

func aToi(str string) (int, error) {
	if str == "" {
		return 0, errors.New("input invalid")
	}
	minus := false
	if []rune(str)[0] == '+' {
		minus = false
		str = string([]rune(str)[1:])
	} else if []rune(str)[0] == '-' {
		minus = true
		str = string([]rune(str)[1:])
	}
	if str == "" {
		return 0, errors.New("input invalid")
	}
	return aToiCore(str, minus)
}

func aToiCore(str string, minus bool) (int, error) {
	num := 0
	flag := 1
	if minus {
		flag = -1
	}
	for _, char := range []rune(str) {
		if char > '0' && char < '9' {
			num = num*10 + flag*int(char-'0')

			if !minus && num > 0x7fffffffffffffff || minus && num < -0x8000000000000000 {
				return 0, errors.New("int flow")
			}
		} else {
			return 0, errors.New("input invalid")
		}
	}
	return num, nil
}
