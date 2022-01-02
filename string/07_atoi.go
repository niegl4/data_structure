package string

import "errors"

/*
六十七
string转int。
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
	for _, char := range []rune(str) {
		if char > '0' && char < '9' {
			flag := 1
			if minus {
				flag = -1
			}
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
