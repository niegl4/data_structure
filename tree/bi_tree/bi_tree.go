package bi_tree

import (
	"data_structure/tree/common"
	"fmt"
)

var (
	BiTreeLeaf1 = BiTree{Data: "H"}
	BiTreeLeaf2 = BiTree{Data: "I"}
	BiTreeLeaf3 = BiTree{Data: "J"}
	BiTreeLeaf4 = BiTree{Data: "K"}
	BiTreeLeaf5 = BiTree{Data: "L"}
	BiTreeLeaf6 = BiTree{Data: "M"}
	BiTreeLeaf7 = BiTree{Data: "N"}
	BiTreeLeaf8 = BiTree{Data: "O"}
	BiTreeNode1 = BiTree{Data: "B", lChild: &BiTreeNode3, rChild: &BiTreeNode4}
	BiTreeNode2 = BiTree{Data: "C", lChild: &BiTreeNode5, rChild: &BiTreeNode6}
	BiTreeNode3 = BiTree{Data: "D", lChild: &BiTreeLeaf1, rChild: &BiTreeLeaf2}
	BiTreeNode4 = BiTree{Data: "E", lChild: &BiTreeLeaf3, rChild: &BiTreeLeaf4}
	BiTreeNode5 = BiTree{Data: "F", lChild: &BiTreeLeaf5, rChild: &BiTreeLeaf6}
	BiTreeNode6 = BiTree{Data: "G", lChild: &BiTreeLeaf7, rChild: &BiTreeLeaf8}
	BiTreeRoot  = BiTree{Data: "A", lChild: &BiTreeNode1, rChild: &BiTreeNode2}
)

type BiTree struct {
	Data   interface{}
	lChild *BiTree
	rChild *BiTree
}

func (bi *BiTree) Operate(data interface{}) {
	fmt.Println(data)
}

// 二叉树的前序遍历
func (bi *BiTree) PreOrderTraverse(op common.Operate) {
	if bi == nil {
		return
	}
	op(bi.Data)
	bi.lChild.PreOrderTraverse(op)
	bi.rChild.PreOrderTraverse(op)
}

// 二叉树的中序遍历
func (bi *BiTree) InOrderTraverse(op common.Operate) {
	if bi == nil {
		return
	}
	bi.lChild.InOrderTraverse(op)
	op(bi.Data)
	bi.rChild.InOrderTraverse(op)
}

// 二叉树的后序遍历
func (bi *BiTree) PostOrderTraverse(op common.Operate) {
	if bi == nil {
		return
	}
	bi.lChild.PostOrderTraverse(op)
	bi.rChild.PostOrderTraverse(op)
	op(bi.Data)
}