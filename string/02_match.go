package string

/*
十九
实现一个函数，用来匹配包含.和*的正则表达式。.表示任意一个字符；*表示它前面的字符可以出现任意次（包括0）。
 */
func match(str, pattern string) bool {
	if str == "" || pattern == "" {
		return false
	}
	strLen := len(str)
	patLen := len(pattern)
	var (
		strSli = make([]rune, 0, len(str))
		patSli = make([]rune, 0, patLen)
	)
	for i := 0; i < strLen; i++ {
		strSli = append(strSli, rune(str[i]))
	}
	for i := 0; i < patLen; i++ {
		patSli = append(patSli, rune(pattern[i]))
	}
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
	if patIdx < patMaxIdx && pat[patIdx+1] == '*' {
		if strIdx < strMaxIdx+1 && (pat[patIdx] == str[strIdx] ||  pat[patIdx] == '.') {
			return matchCore(str, pat, strIdx+1, strMaxIdx, patIdx+2, patMaxIdx) || //a【.】..	a*【.】..
				matchCore(str, pat, strIdx+1, strMaxIdx, patIdx, patMaxIdx) || //a【.】.. 	【a】...
				matchCore(str, pat, strIdx, strMaxIdx, patIdx+2, patMaxIdx) // 【a】... 		a*【.】..
		} else {
			return matchCore(str, pat, strIdx, strMaxIdx, patIdx+2, patMaxIdx)
		}
	}
	if strIdx < strMaxIdx+1 && (str[strIdx] == pat[patIdx] || pat[patIdx] == '.') {
		return matchCore(str, pat, strIdx+1, strMaxIdx, patIdx+1, patMaxIdx)
	}
	return false
}