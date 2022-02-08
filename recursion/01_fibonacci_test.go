package recursion

import "testing"

func TestFibonacci(t *testing.T) {
	t.Log(fibonacci(6))
	t.Log(fibonacci(0x7fffffffffffffff))

	t.Log(0x7fffffffffffffff)
	t.Log(-0x8000000000000000)
}
