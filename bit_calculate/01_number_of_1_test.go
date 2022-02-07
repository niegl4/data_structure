package bit_calculate

import "testing"

func TestNumberOf1(t *testing.T) {
	t.Log(numberOf1Version1(0b1001))
	t.Log(numberOf1Version1(0x7fffffffffffffff)) //最大整数

	t.Log(numberOf1Version1(-1)) //最大负数-1，它的二进制表示中，64位全是1
	//t.Log(numberOf1Version1(0xffffffffffffffff))//todo：为什么0xffffffffffffffff这条测试用例通不过呢？
	//t.Log(numberOf1Version1(0x8fffffffffffffff))//todo：为什么这个表示最小负数的用例，通不过呢？
	t.Log(numberOf1Version1(-9223372036854775808)) //todo：为什么这个表示最小负数的用例，能通过呢？

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
