package _2_bi_search_tree

import (
	"data_structure/tree/common"
)

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
	{
		node6.parent = &node10
		node6.lChild = &node1
		node6.rChild = &node2
	}
	{
		node7.parent = &node12
		node7.rChild = &node3
	}
	node8.parent = &node12

	{
		node9.parent = &node13
		node9.rChild = &node4
	}
	{
		node10.parent = &node13
		node10.lChild = &node5
		node10.rChild = &node6
	}
	node11.parent = &node14
	{
		node12.parent = &node14
		node12.lChild = &node7
		node12.rChild = &node8
	}

	{
		node13.parent = &root
		node13.lChild = &node9
		node13.rChild = &node10
	}
	{
		node14.parent = &root
		node14.lChild = &node11
		node14.rChild = &node12
	}

	{
		root.parent = nil
		root.lChild = &node13
		root.rChild = &node14
	}
}

//todo：暂时先不考虑有多个相等data的节点
type BiSearchTree struct {
	Data   int
	parent *BiSearchTree
	lChild *BiSearchTree
	rChild *BiSearchTree
}

//增加，返回新增加的节点
func (biS *BiSearchTree) Insert(data int) (newNode *BiSearchTree) {
	//空树，直接返回nil
	if biS == nil {
		return nil
	}
	if data < biS.Data {
		if biS.lChild == nil {
			biS.lChild = &BiSearchTree{Data: data}
			biS.lChild.parent = biS
			return biS.lChild
		} else {
			return biS.lChild.Insert(data)
		}
	} else {
		if biS.rChild == nil {
			biS.rChild = &BiSearchTree{Data: data}
			biS.rChild.parent = biS
			return biS.rChild
		} else {
			return biS.rChild.Insert(data)
		}
	}
}

//删除，返回被删除节点位置上的新节点。
//没有该节点，status为-1；
//删除该节点后为空树，status为0；
//删除该节点后那个位置为nil，但不是空树，status为1；
//删除该节点后那个位置不为nil，status为2.
func (biS *BiSearchTree) Remove(data int) (newNode *BiSearchTree, status int) {
	node := biS.SearchReturnNode(data)
	if node == nil {
		return nil, -1
	}
	//查找
	defer func() {node=nil}()
	nodeParent := node.parent
	nodeLChild := node.lChild
	nodeRChild := node.rChild

	//1.要删除的节点，是叶子节点
	if nodeLChild == nil && nodeRChild == nil {
		//要删除节点就是根节点
		if nodeParent == nil {
			return nil, 0
		}
		if node.Data < nodeParent.Data {
			nodeParent.lChild = nil
		} else {
			nodeParent.rChild = nil
		}
		return nil, 1
	}

	//2.要删除的节点，只有一个子节点
	if nodeLChild == nil && nodeRChild != nil || nodeLChild != nil && nodeRChild == nil {
		tmp := &BiSearchTree{}
		if nodeLChild == nil && nodeRChild != nil {
			tmp = nodeRChild
		}
		if nodeLChild != nil && nodeRChild == nil {
			tmp = nodeLChild
		}
		//要删除节点就是根节点
		if nodeParent == nil {
			tmp.parent = nil
			return tmp, 2
		}
		if tmp.Data < nodeParent.Data {
			nodeParent.lChild = tmp
		} else {
			nodeParent.rChild = tmp
		}
		tmp.parent = nodeParent
		return tmp, 2
	}

	//3.要删除的节点，有两个子节点
	minNode := nodeRChild.findMinNode()
	//最小节点与lChild节点，建立联系
	minNode.lChild = nodeLChild
	nodeLChild.parent = minNode

	//最小节点与父节点，建立联系
	//在minNode修改parent之前，先保存旧的parent节点
	minNodeOldParent := minNode.parent
	if nodeParent == nil {
		minNode.parent = nil
	} else {
		minNode.parent = nodeParent
		if node.Data < nodeParent.Data {
			nodeParent.lChild = minNode
		} else {
			nodeParent.rChild = minNode
		}
	}

	//最小节点与rChild节点，建立联系。最小节点，不是要删除节点的rChild；反之，不用做任何处理。
	if minNode != node.rChild {
		minNode.rChild = nodeRChild
		nodeRChild.parent = minNode
		//旧的parent节点的lChild，要赋值nil
		minNodeOldParent.lChild = nil
	}

	return minNode, 2
}

//寻找最小节点
func (biS *BiSearchTree) findMinNode() (minNode *BiSearchTree) {
	if biS == nil {
		return nil
	}
	if biS.lChild == nil {
		return biS
	} else {
		return biS.lChild.findMinNode()
	}
}

//删除，返回新插入的节点
func (biS *BiSearchTree) Change(srcData, dstData int) *BiSearchTree {
	if biS == nil {
		return nil
	}
	newNode, status := biS.Remove(srcData)
	if status == -1 {
		return nil
	} else if status == 0 {
		root = BiSearchTree{Data: dstData}
		return &root
	} else if status == 1 {
		return root.Insert(dstData)
	} else {
		//根节点变更
		if newNode.parent == nil {
			root = *newNode
		}
		return root.Insert(dstData)
	}
}

//搜索，返回该节点
func (biS *BiSearchTree) SearchReturnNode(data int) (node *BiSearchTree) {
	//空树，直接返回nil
	if biS == nil {
		return nil
	}
	//是叶子节点，仍然没找到
	if biS.lChild == nil && biS.rChild == nil && biS.Data != data {
		return nil
	}

	if biS.Data == data {
		return biS
	} else {
		if data < biS.Data {
			return biS.lChild.SearchReturnNode(data)
		} else {
			return biS.rChild.SearchReturnNode(data)
		}
	}
}

//搜索，返回查询路径以及是否找到。[此处只为了演练：返回结果包含逐级递归结果]
func (biS *BiSearchTree) SearchReturnRoute(data int) (route []int, isFind bool) {
	//空树，直接返回nil
	if biS == nil {
		return nil, false
	}
	//是叶子节点，仍然没找到
	if biS.lChild == nil && biS.rChild == nil && biS.Data != data {
		return []int{biS.Data}, false
	}

	//找到，返回该节点的数据，true
	if biS.Data == data {
		return []int{biS.Data}, true
	} else {
		//localRoute包含本节点的数据；tmpRoute接收下一级的route;下一级的isFind，就代表了本级的isFind
		var (
			localRoute = []int{biS.Data}
			tmpRoute   []int
			tmpIsFind  bool
		)
		if data < biS.Data {
			tmpRoute, tmpIsFind = biS.lChild.SearchReturnRoute(data)
		} else {
			tmpRoute, tmpIsFind = biS.rChild.SearchReturnRoute(data)
		}
		return append(localRoute, tmpRoute...), tmpIsFind
	}
}

//二插搜索树的中序遍历，结果是有序数组
func (biS *BiSearchTree) inOrderTraverse(op common.Operate) {
	if biS == nil {
		return
	}
	biS.lChild.inOrderTraverse(op)
	op(biS.Data)
	biS.rChild.inOrderTraverse(op)
}
