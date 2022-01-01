package string

/*
五十八-1
翻转字符串。
*/
func reverseSentence(str string) string {
	if str == "" {
		return ""
	}

	bytes := []byte(str)
	length := len(bytes)
	start := 0
	end := length - 1
	reverse(bytes, start, end)

	for i := 0; i < length; i++ {
		if i == 0 && bytes[0] != ' ' {
			j := i + 1
			for ; j < length; j++ {
				if bytes[j] == ' ' {
					break
				}
			}
			reverse(bytes, i, j-1)
		} else {
			if bytes[i] == ' ' {
				j := i + 1
				for ; j < length; j++ {
					if bytes[j] == ' ' {
						break
					}
				}
				reverse(bytes, i+1, j-1)
			}
		}
	}
	return string(bytes)
}

func reverse(bytes []byte, start, end int) {
	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
}

/*
五十八-2
左旋转字符串。
*/
func leftRotateStr(str string, k int) string {
	length := len(str)
	if length == 0 || k >= length {
		return str
	}

	bytes := []byte(str)
	part1Start := 0
	part1End := k - 1
	part2Start := k
	part2End := length - 1

	reverse(bytes, part1Start, part1End)
	reverse(bytes, part2Start, part2End)
	reverse(bytes, part1Start, part2End)
	return string(bytes)
}
