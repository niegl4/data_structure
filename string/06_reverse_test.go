package string

import "testing"

func TestReverseSentence(t *testing.T) {
	t.Logf("%#v", reverseSentence("I am a student."))
	t.Logf("%#v", reverseSentence("I am a student. "))
	t.Logf("%#v", reverseSentence(" I am a student."))
	t.Logf("%#v", reverseSentence(" I am a student. "))
	t.Logf("%#v", reverseSentence("  I              am a student. "))
}

func TestLeftRotateStr(t *testing.T) {
	t.Logf("%#v", leftRotateStr("abcdefg", 2))
}
