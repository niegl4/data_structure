package dynamic_algorithm

import "testing"

func TestPackageByBacktrack(t *testing.T) {
	objs := []int{2, 2, 4, 6, 3}
	t.Log(packageByBacktrack(objs, 9))
	t.Log(packageWithMatrix(objs, 9))
	t.Log(packageWithArray(objs, 9))

	values := []int{3, 4, 8, 9, 6}
	t.Log(packageValueWithArray(objs, values, 9))
}
