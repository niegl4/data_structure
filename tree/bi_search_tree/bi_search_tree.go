package bi_search_tree

import "data_structure/tree/common"

var (
	/*
		         33
		   16          50
		13    18 	34	  58
		 15 17  25      51  66
		       19 27     55
	*/
	node1 = BiSearchTree{Data: 19}
	node2 = BiSearchTree{Data: 27}
	node3 = BiSearchTree{Data: 55}

	node4 = BiSearchTree{Data: 15}
	node5 = BiSearchTree{Data: 17}
	node6 = BiSearchTree{Data: 25}
	node7 = BiSearchTree{Data: 51}
	node8 = BiSearchTree{Data: 66}

	node9  = BiSearchTree{Data: 13}
	node10 = BiSearchTree{Data: 18}
	node11 = BiSearchTree{Data: 34}
	node12 = BiSearchTree{Data: 58}

	node13 = BiSearchTree{Data: 16}
	node14 = BiSearchTree{Data: 50}

	root = BiSearchTree{Data: 33}
)

func init() {
	node1.parent = &node6
	node2.parent = &node6
	node3.parent = &node7

	node4.parent = &node9
	node5.parent = &node10
	node6.parent = &node10
	node6.lChild = &node1
	node6.rChild = &node2
	node7.parent = &node12
	node7.rChild = &node3
	node8.parent = &node12

	node9.parent = &node13
	node9.rChild = &node4
	node10.parent = &node13
	node10.lChild = &node5
	node10.rChild = &node6
	node11.parent = &node14
	node12.parent = &node14
	node12.lChild = &node7
	node12.rChild = &node8

	node13.parent = &root
	node13.lChild = &node9
	node13.rChild = &node10
	node14.parent = &root
	node14.lChild = &node11
	node14.rChild = &node12

	root.parent = nil
	root.lChild = &node13
	root.rChild = &node14
}

type BiSearchTree struct {
	Data   int
	parent *BiSearchTree
	lChild *BiSearchTree
	rChild *BiSearchTree
}

func (biS *BiSearchTree) Insert(data int) {

}

func (biS *BiSearchTree) Remove(data int) {

}

func (biS *BiSearchTree) Change(srcData, dstData int) {

}

func (biS *BiSearchTree) Search(data int) (route []int, isFind bool) {
	//空树，直接返回nil
	if biS == nil {
		return nil, false
	}
	//是叶子节点，仍然没找到
	if biS.lChild == nil && biS.rChild == nil && biS.Data != data {
		return []int{biS.Data}, false
	}

	if biS.Data == data {
		return []int{biS.Data}, true
	} else {
		var (
			localRoute = []int{biS.Data}
			tmpRoute   []int
			tmpIsFind  bool
		)
		if data < biS.Data {
			tmpRoute, tmpIsFind = biS.lChild.Search(data)
		} else {
			tmpRoute, tmpIsFind = biS.rChild.Search(data)
		}
		return append(localRoute, tmpRoute...), tmpIsFind
	}
}

func (biS *BiSearchTree) inOrderTraverse(op common.Operate) {
	if biS == nil {
		return
	}
	biS.lChild.inOrderTraverse(op)
	op(biS.Data)
	biS.rChild.inOrderTraverse(op)
}
