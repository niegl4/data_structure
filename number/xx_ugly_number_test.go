package number

import "testing"

func TestGetUglyNumber(t *testing.T) {
	t.Log(getUglyNumberV1(1500)) //859963392
	t.Log(getUglyNumberV2(1500)) //859963392
}
