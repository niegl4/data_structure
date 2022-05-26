package number

import (
	"errors"
)

/*
十六
实现函数func power(base float64, exponent int) (float64,error)
不考虑大数问题
*/

func power(base float64, exponent int) (float64, error) {
	if base == 0 && exponent < 0 {
		return 0, errors.New("invalid input")
	}

	exp := exponent
	if exponent < 0 {
		exp = -exponent
	}

	result := calPowerV1(base, uint64(exp))

	if exponent < 0 {
		result = 1 / result
	}
	return result, nil
}

//v1中，O(exponent)
func calPowerV1(base float64, exponent uint64) float64 {
	result := float64(0)
	for i := uint64(1); i <= exponent; i++ {
		result *= base
	}
	return result
}

//v2中，O(log(exponent))。
//右移代替除以2
//位与运算代替求余判断一个数是奇数还是偶数
func calPowerV2(base float64, exponent uint64) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}

	result := calPowerV2(base, exponent>>1)
	result *= result

	if exponent&0x1 == 1 {
		result *= base
	}
	return result
}
