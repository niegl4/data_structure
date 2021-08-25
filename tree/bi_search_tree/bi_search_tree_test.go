package bi_search_tree

import "testing"

func TestBiSearchTree_Search(t *testing.T) {
	route, isFind := root.Search(100)
	t.Log(isFind, route)
	route, isFind = root.Search(27)
	t.Log(isFind, route)
}
