package _1_bi_tree

import (
	"data_structure/tree/common"
	"fmt"
)

var (
	/*
	       A
	   B       C
	 D   E   F   G
	H I J K L M N O
	 */
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

// 二叉树的层序遍历
func (bi *BiTree) LevelOrderTraverse(op common.Operate) {
	if bi == nil {
		return
	}
	levelSet := make([][]interface{}, 0)
	queue := []*BiTree{bi}
	for len(queue) > 0 {
		//[当前]队列长度，在队列中插入数据前计算一次，内层循环入队的元素不影响n
		n := len(queue)
		level := make([]interface{}, 0)
		//内层循环相当于处理同一层的逻辑
		for i := 0; i < n; i++ {
			//1.子节点入队
			if queue[i].lChild != nil {
				queue = append(queue, queue[i].lChild)
			}
			if queue[i].rChild != nil {
				queue = append(queue, queue[i].rChild)
			}
			//2.本节点的值入本层结果
			level = append(level, queue[i].Data)
		}
		levelSet = append(levelSet, level)
		//queue截取掉本层
		queue = queue[n:]
	}
	op(levelSet)
}