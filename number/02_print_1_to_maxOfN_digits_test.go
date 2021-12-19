package number

import (
	"testing"
)

func TestPrint1ToMaxOfDigits(t *testing.T) {
	//t.Log('9' - '0') //字符之间可以运算
	//t.Log(9 - '0') //不可控
	//t.Log(48 - '0') //48在字符中表达为'0'

	print1ToMaxOfDigits(1)

	print1ToMaxOfDigitsV2(3)
}
