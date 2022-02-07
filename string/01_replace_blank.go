package string

/*
五-1
把字符串中的每个空格替换成%20
*/
func replaceBlank(str string) (newStr string) {
	if str == "" {
		return
	}

	blankNum := 0
	oriSli := make([]rune, 0)
	for _, char := range str {
		if char == ' ' {
			blankNum++
		}
		oriSli = append(oriSli, char)
	}

	oriIndex := len(oriSli) - 1
	newIndex := oriIndex + 2*blankNum
	newStrSli := make([]rune, newIndex+1)

	for oriIndex >= 0 && newIndex >= oriIndex {
		if oriSli[oriIndex] == ' ' {
			newStrSli[newIndex] = '0'
			newIndex--
			newStrSli[newIndex] = '2'
			newIndex--
			newStrSli[newIndex] = '%'
		} else {
			newStrSli[newIndex] = oriSli[oriIndex]
		}
		oriIndex--
		newIndex--
	}

	return string(newStrSli)
}
