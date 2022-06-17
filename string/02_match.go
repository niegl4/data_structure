package string

/*
十九
实现一个函数，用来匹配包含.和*的正则表达式。.表示任意一个字符；*表示它前面的字符可以出现任意次（包括0）。
*
*/
func match(str, pattern string) bool {
	if str == "" || pattern == "" {
		return false
	}
	strLen := len(str)
	patLen := len(pattern)
	strSli := []rune(str)
	patSli := []rune(pattern)
	return matchCore(strSli, patSli, 0, strLen-1, 0, patLen-1)
}

func matchCore(str, pat []rune, strIdx, strMaxIdx, patIdx, patMaxIdx int) bool {
	//都是末尾
	if strIdx == strMaxIdx+1 && patIdx == patMaxIdx+1 {
		return true
	}
	//str不是末尾，pat是末尾
	if strIdx < strMaxIdx+1 && patIdx == patMaxIdx+1 {
		return false
	}

	//str末尾，pat不是末尾，不确定：“a”,"a*"。strIdx == strMaxIdx+1
	//str不是末尾，pat不是末尾，不确定。
	//当前第二个字符是*
	if patIdx < patMaxIdx && pat[patIdx+1] == '*' {
		if strIdx < strMaxIdx+1 && (pat[patIdx] == str[strIdx] || pat[patIdx] == '.') { //当前匹配
			return matchCore(str, pat, strIdx+1, strMaxIdx, patIdx+2, patMaxIdx) || //str步进1，pattern步进2
				matchCore(str, pat, strIdx+1, strMaxIdx, patIdx, patMaxIdx) || //str步进1，pattern不动
				matchCore(str, pat, strIdx, strMaxIdx, patIdx+2, patMaxIdx) // str不动，pattern步进2
		} else { //当前不匹配
			return matchCore(str, pat, strIdx, strMaxIdx, patIdx+2, patMaxIdx) // str不动，pattern步进2
		}
	}
	//当前第二个字符不是*，当前第一个字符匹配，接着匹配后面的字符
	if strIdx < strMaxIdx+1 && (str[strIdx] == pat[patIdx] || pat[patIdx] == '.') {
		return matchCore(str, pat, strIdx+1, strMaxIdx, patIdx+1, patMaxIdx) //str步进1，pattern步进1
	}
	return false
}
