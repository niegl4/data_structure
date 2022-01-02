package bit_calculate

import "testing"

func TestAdd(t *testing.T) {
	t.Log(add(5, 17))
}

func TestSwap(t *testing.T) {
	t.Log(swapV1(1, 2))
	t.Log(swapV2(1, 2))
}
