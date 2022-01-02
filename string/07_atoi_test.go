package string

import "testing"

func TestAToi(t *testing.T) {
	//Range: -9223372036854775808 through 9223372036854775807.
	t.Log(0x7fffffffffffffff)
	t.Log(-0x8000000000000000)
	t.Log(aToi("+123"))
}
