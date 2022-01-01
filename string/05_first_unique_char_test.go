package string

import (
	"testing"
)

func TestFirstUniqueChar(t *testing.T) {
	t.Log(firstUniqueCharV1("azbaccdeff"))
	t.Log(firstUniqueCharV2("我a是中国人我"))
}

func TestFirstUniqueInStream(t *testing.T) {
	charChan := make(chan byte, 0)
	go func() {
		str := "abaccbeffmnkjgln"
		for _, s := range str {
			charChan <- byte(s)
		}
	}()

	firstUniqueInStream(charChan)
}
