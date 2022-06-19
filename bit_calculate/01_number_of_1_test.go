package bit_calculate

import (
	"math"
	"testing"
)

func TestNumberOf1(t *testing.T) {
	t.Log(numberOf1Version1(0b1001))
	t.Log(numberOf1Version1(0x7fffffffffffffff)) //最大整数

	t.Log(numberOf1Version1(-1)) //最大负数-1，它的二进制表示中，64位全是1
	//t.Log(numberOf1Version1(-0xffffffffffffffff))  //todo：最大负数的十六进制不是0xffff ffff ffff ffff？
	t.Log(numberOf1Version1(math.MinInt64))        //最小负数
	t.Log(numberOf1Version1(-0x8000000000000000))  //最小负数的十六进制表达
	t.Log(numberOf1Version1(-9223372036854775808)) //最小负数的十进制表达

	t.Log(numberOf1Version2(0b1001))
	t.Log(numberOf1Version2(0x7fffffffffffffff)) //最大整数

	t.Log(numberOf1Version2(-1))                   //最大负数-1，它的二进制表示中，64位全是1
	t.Log(numberOf1Version2(-9223372036854775808)) //最小整数
}

func TestJudgeNumber(t *testing.T) {
	t.Log(judgeNumber(1))
	t.Log(judgeNumber(2))
	t.Log(judgeNumber(4))
	t.Log(judgeNumber(5))
}

func TestChangeMtoN(t *testing.T) {
	t.Log(changeMtoN(0b1010, 0b1101))
}
