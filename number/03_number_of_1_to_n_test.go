package number

import "testing"

func TestNumberOf1ToN(t *testing.T) {
	/*
			1~45 (递归位（1）+ 最高位（10^1）+ 其他位（4*1*10^0）=15)
		1~5		6~45

			1~345(递归位（15）+ 最高位（10^2）+ 其他位（3*2*10^1）= 175)
		1~45	46~345
	*/
	//45: 最高位
	t.Log(numberOf1ToN(345)) // todo:21345 => 10000 + 2*4*1000 + 1345中的1(346+ )
}
