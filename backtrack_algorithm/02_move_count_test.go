package backtrack_algorithm

import "testing"

func TestMoveCount(t *testing.T) {
	t.Log(moveCount(4, 4, 4))   //13
	t.Log(moveCount(4, 11, 11)) //15
}
