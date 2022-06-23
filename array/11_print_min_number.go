package array

import (
	"fmt"
	"sort"
	"strings"
)

/*
四十五
输入一个正整数数组，把数组里所有数字拼接起来排成一个数，输出所有能拼接的数字中最小的那个。

设计一种排序规则，外加大数问题
*/

func printMinNumber(arr []int) string {
	if len(arr) <= 0 {
		return ""
	}

	var arrOfStr []string
	for _, b := range arr {
		str := fmt.Sprintf("%d", b)
		arrOfStr = append(arrOfStr, str)
	}
	sort.Slice(arrOfStr, func(i, j int) bool {
		str1 := arrOfStr[i] + arrOfStr[j]
		str2 := arrOfStr[j] + arrOfStr[i]
		return str1 < str2
	})

	return strings.Join(arrOfStr, "")
}
