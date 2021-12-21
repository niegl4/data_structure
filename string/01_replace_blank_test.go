package string

import "testing"

func TestReplaceBlank(t *testing.T) {
	testSet := [][]rune{
		{'我', ' ', '是', ' ', '学', '生'},
		{' ', '我', ' ', '是', ' ', '学', '生'},
		{'我', ' ', '是', ' ', '学', '生', ' '},
		{' '},
	}
	for _, s := range testSet{
		replaceBlank(&s)
		t.Log(string(s))
	}
}
