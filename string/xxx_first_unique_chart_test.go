package string

import (
	"testing"
)

func TestFirstUniqueChar(t *testing.T) {
	t.Log(firstUniqueCharV1("azbaccdeff"))
	t.Log(firstUniqueCharV2("我a是中国人我"))
}
