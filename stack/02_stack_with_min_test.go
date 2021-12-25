package stack

import "testing"

func TestStackWithMin_Min(t *testing.T) {
	//var stack *stackWithMin
	//val, err := stack.min()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(val)

	stack := NewStackWithMin()
	var err error
	if err = stack.push(3); err != nil {
		t.Fatal(err)
	}
	if err = stack.push(4); err != nil {
		t.Fatal(err)
	}
	if err = stack.push(2); err != nil {
		t.Fatal(err)
	}
	if err = stack.push(1); err != nil {
		t.Fatal(err)
	}

	minVal, err := stack.min()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("cur min: %d", minVal)

	for i := 0; i < 10; i++ {
		data, err := stack.pop()
		if err != nil {
			t.Fatal(err)
		}
		minVal, err = stack.min()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("data:%d, min:%d", data, minVal)
	}
}
