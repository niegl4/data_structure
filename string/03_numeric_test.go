package string

import "testing"

func TestIsNumeric(t *testing.T) {
	strSet := []string{
		"+100",
		"5e2",
		"-123",
		"3.1416",
		"-1E-16",

		"12e",
		"1a3.14",
		"1.2.3",
		"+-5",
		"12e+5.4",
	}
	for _, str := range strSet {
		t.Log(isNumeric(str))
	}
}
