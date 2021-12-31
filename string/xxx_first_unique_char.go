package string

import "fmt"

/*
五十-1
字符串中第一个只出现一次的字符。
 */
func firstUniqueCharV1(str string) string {
	if str == "" {
		return ""
	}

	charHashTable := make([]int, 256)
	for _, b := range []byte(str) {
		charHashTable[b]++
	}

	for _, b := range []byte(str) {
		if charHashTable[b] == 1 {
			return fmt.Sprintf("%c", b)
		}
	}
	return ""
}

//支持汉字
func firstUniqueCharV2(str string) string {
	if str == "" {
		return ""
	}

	charHashTable := make([]int, 4294967296)
	for _, b := range []rune(str) {
		if b < 0 {
			b = int32(int(b) + 4294967296)
		}
		charHashTable[b]++
	}

	for _, b := range []rune(str) {
		if charHashTable[b] == 1 {
			if b > 2147483647 {
				b = int32(int(b) - 4294967296)
			}
			return fmt.Sprintf("%c", b)
		}
	}
	return ""
}

