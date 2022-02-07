package string

import "testing"

func TestReplaceBlank(t *testing.T) {
	testSet := []string{
		"我 是 学生",
		" 我 是 学生",
		"我 是 学生 ",
		" ",
	}
	for _, s := range testSet {
		t.Log(replaceBlank(s))
	}
}
