package string

import "testing"

func TestMatch(t *testing.T) {
	strSet := [][2]string{
		{"aaa", "a.a"},
		{"aaa", "ab*ac*a"},
		{"aaa", "ab.a"},
		{"aaa", "ab*a"},
	}
	for _, set := range strSet {
		t.Log(match(set[0], set[1]))
	}
}
