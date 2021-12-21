package string

/*
五-1
把字符串中的每个空格替换成%20
 */
func replaceBlank(str *[]rune) {
	if str == nil {
		return
	}

	blankNum := 0
	for _, char := range *str {
		if char == ' ' {
			blankNum++
		}
	}

	oriIndex := len(*str) - 1
	newIndex := oriIndex + 2 * blankNum
	//数组扩充
	for i := 0; i < blankNum; i++ {
		*str = append(*str, ' ', ' ')
	}

	for oriIndex >= 0 && newIndex >= oriIndex {
		if (*str)[oriIndex] == ' ' {
			(*str)[newIndex] = '0'
			newIndex--
			(*str)[newIndex] = '2'
			newIndex--
			(*str)[newIndex] = '%'
		} else {
			(*str)[newIndex] = (*str)[oriIndex]
		}
		oriIndex--
		newIndex--
	}

	return
}
