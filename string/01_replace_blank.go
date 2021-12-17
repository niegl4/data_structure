package string

/*
五
把字符串中的每个空格替换成%20
go中不支持字符串的索引赋值。。。
 */
//func replaceBlank(str *string) {
//	if str == nil {
//		return
//	}
//
//	blankNum := 0
//	for _, char := range *str {
//		if char == ' ' {
//			blankNum++
//		}
//	}
//
//	a := "abc"
//	bl := a[2] == 'd'
//
//	oriIndex := len(*str) - 1
//	newIndex := oriIndex + 2 * blankNum
//	for oriIndex >= 0 && newIndex >= oriIndex {
//		if (*str)[oriIndex] == ' ' {
//			(*str)[newIndex] = '0'
//			*str[newIndex] = '0'
//
//		}
//	}
//
//	return
//}
