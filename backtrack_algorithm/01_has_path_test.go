package backtrack_algorithm

import "testing"

func TestHasPath(t *testing.T) {
	matrix := [][]byte{
		{'b', 'b', 't', 'g'},
		{'c', 'f', 'c', 's'},
		{'j', 'd', 'e', 'h'},
	}
	t.Log(hasPath(matrix, "bfce"))
	t.Log(hasPath(matrix, "abfb"))
}
