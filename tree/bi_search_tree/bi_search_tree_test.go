package bi_search_tree

import (
	"data_structure/tree/common"
	"testing"
)

func TestBiSearchTree_Insert(t *testing.T) {
	newNode := root.Insert(20)
	t.Log(newNode.Data)
	TestBiSearchTree_inOrderTraverse(t)
}

func TestBiSearchTree_Remove(t *testing.T) {
	tmp := make([]interface{}, 0, 8)
	var op common.Operate = func(data interface{}) {
		tmp = append(tmp, data)
	}
	testFunc := func(t *testing.T, data int) {
		t.Log("-------------------")
		removeNode, status := root.Remove(data)
		//数中不存在该节点
		if status == -1 {
			t.Log("not found")
			t.Log("-------------------")
			return
		}

		defer func() {
			t.Log(tmp)
			tmp = tmp[0:0]
			t.Log("-------------------")
		}()
		//被删除的是叶节点
		if removeNode == nil {
			t.Log(data, "is removed from bi-search-tree")
			root.inOrderTraverse(op)
			return
		} else {
			t.Log("newNode", removeNode.Data)
			if removeNode.parent != nil {
				t.Log("newNode parent", removeNode.parent.Data)
			}
			if removeNode.lChild != nil {
				t.Log("newNode lChild", removeNode.lChild.Data)
			}
			if removeNode.rChild != nil {
				t.Log("newNode rChild", removeNode.rChild.Data)
			}
			if removeNode.parent == nil {//被删除节点的位置，没有父节点，说明被删除节点就是根节点
				removeNode.inOrderTraverse(op)
			} else {//根节点还是旧的根节点
				root.inOrderTraverse(op)
			}
		}
	}

	testFunc(t, 13)
	testFunc(t, 18)
	testFunc(t, 55)
	testFunc(t, 33)
	testFunc(t, 100)
}

func TestBiSearchTree_Change(t *testing.T) {
	tmp := make([]interface{}, 0, 8)
	var op common.Operate = func(data interface{}) {
		tmp = append(tmp, data)
	}
	testFunc := func(t *testing.T, src, dst int) () {
		t.Log("-------------------")
		newNode := root.Change(src, dst)
		if newNode == nil {
			t.Log(src, "not found")
			t.Log("-------------------")
			return
		}

		defer func() {
			t.Log(tmp)
			tmp = tmp[0:0]
			t.Log("-------------------")
		}()
		t.Log(newNode.Data)
		if newNode.parent == nil {
			newNode.inOrderTraverse(op)
		} else {
			root.inOrderTraverse(op)
		}
		return
	}
	//status -1
	testFunc(t, -1, 0)
	//status 1
	testFunc(t, 55, 80)
	//status 2 不变更根节点
	testFunc(t, 58, 100)
	//status 2 变更根节点
	testFunc(t, 33, 11)
}

func TestBiSearchTree_SearchReturnNode(t *testing.T) {
	node := root.SearchReturnNode(100)
	t.Log(node)
	node = root.SearchReturnNode(27)
	t.Log(node)
}

func TestBiSearchTree_SearchReturnRoute(t *testing.T) {
	route, isFind := root.SearchReturnRoute(100)
	t.Log(isFind, route)
	route, isFind = root.SearchReturnRoute(27)
	t.Log(isFind, route)
}

func TestBiSearchTree_inOrderTraverse(t *testing.T) {
	tmp := make([]interface{}, 0, 8)
	var op common.Operate = func(data interface{}) {
		tmp = append(tmp, data)
	}
	root.inOrderTraverse(op)
	t.Log(tmp)
}
