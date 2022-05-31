package string

import "fmt"

/*
五十-1
字符串中第一个只出现一次的字符。
哈希表中的值表示：该字符出现的次数。
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

/*
五十-2
字符流中，第一个只出现一次的字符。

唯一难点就在于：设置哈希表的值。>=0：该字符第一次出现的索引；-1：该字符从未出现；-2：该字符已经重复出现
*/
func firstUniqueInStream(strChain <-chan byte) {
	idx := 0
	str := make([]byte, 0, 100)
	charHashTable := make([]int, 256)
	for idx := range charHashTable {
		charHashTable[idx] = -1
	}

	for {
		b := <-strChain
		str = append(str, b)

		if charHashTable[b] >= 0 {
			charHashTable[b] = -2
		} else if charHashTable[b] == -1 {
			charHashTable[b] = idx
		}

		//通过遍历string，判断当前第一个只出现一次的char。其实最好改成遍历哈希表获取该字符。
		minIdx := -1
		var c byte
		for _, char := range str {
			if charHashTable[char] >= 0 {
				if minIdx < 0 {
					minIdx = charHashTable[char]
					c = char
				} else {
					if charHashTable[char] < minIdx {
						minIdx = charHashTable[char]
						c = char
					}
				}
			}
		}

		fmt.Printf("idx:%d, char:%c\n", idx, c)
		idx++

		if idx >= 10 {
			return
		}
	}
}
