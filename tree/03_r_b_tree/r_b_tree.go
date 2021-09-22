package _3_r_b_tree

import (
	"data_structure/tree/common"
	"fmt"
	"strings"
)

var count int

type TreeNode struct {
	data   float64
	color  string //比二叉查找树要多出一个颜色属性
	lchild *TreeNode
	rchild *TreeNode
	parent *TreeNode
}

type RBTree struct {
	root   *TreeNode
	cur    *TreeNode
	create *TreeNode
}

func (rbt *RBTree) Add(data float64) {
	rbt.create = new(TreeNode)
	rbt.create.data = data
	rbt.create.color = "red"

	if !rbt.IsEmpty() {
		rbt.cur = rbt.root
		for {
			if data < rbt.cur.data {
				//如果要插入的值比当前节点的值小，则当前节点指向当前节点的左孩子，如果
				//左孩子为空，就在这个左孩子上插入新值
				if rbt.cur.lchild == nil {
					rbt.cur.lchild = rbt.create
					rbt.create.parent = rbt.cur
					break
				} else {
					rbt.cur = rbt.cur.lchild
				}

			} else if data > rbt.cur.data {
				//如果要插入的值比当前节点的值大，则当前节点指向当前节点的右孩子，如果
				//右孩子为空，就在这个右孩子上插入新值
				if rbt.cur.rchild == nil {
					rbt.cur.rchild = rbt.create
					rbt.create.parent = rbt.cur
					break
				} else {
					rbt.cur = rbt.cur.rchild
				}

			} else {
				//如果要插入的值在树中已经存在，则退出
				return
			}
		}

	} else {
		rbt.root = rbt.create
		rbt.root.color = "black"
		rbt.root.parent = nil
		return
	}

	//插入节点后对红黑性质进行修复
	rbt.insertBalanceFixup(rbt.create)
}

func (rbt *RBTree) DeleteByMe(data float64) {
	node := rbt.Search(data)
	if node == nil {
		return
	}
	freeNode := func(n *TreeNode) {
		n.color = "no mean"
		n.lchild = nil
		n.rchild = nil
		n.parent = nil
	}
	switch {
	//被删除节点是叶节点
	case node.lchild == nil && node.rchild == nil:
		if node.parent == nil {
			rbt.root = nil
			return
		}
		//红色节点删除不影响黑色平衡，只用把父节点的指针指向nil即可
		if node.color == "red" {
			if node.parent.lchild == node {
				node.parent.lchild = nil
			} else {
				node.parent.rchild = nil
			}
			freeNode(node)
			return
		} else {
			//被删除节点是叶节点，并且是黑色。人为给它的父节点添加一个特殊的nil节点，该nil节点有两个黑色。保证删除后，黑色平衡
			//这时，它就是删除操作中的2阶段调整。
			trackNode := &TreeNode{
				parent: node.parent,
				color: "black+black",
			}
			if node.parent.lchild == node {
				node.parent.lchild = trackNode
			} else {
				node.parent.rchild = trackNode
			}
			rbt.deleteStage2(trackNode)
		}

	/*
		被删除节点只有一个子节点的时候，被删节点一定是黑，非nil子节点一定是红
			 X-> a●                             b●
				 / \             --------\      / \
			   ○b  nil           --------/    nil  nil
	*/
	case node.lchild != nil && node.rchild == nil:
		//node刚好又是root节点
		if node.parent == nil {
			node.lchild.parent = nil
			node.lchild.color = "black"
			rbt.root = node.lchild
			freeNode(node)
			return
		}
		//node.parent与node的相对位置关系
		if node.parent.lchild == node {
			node.parent.lchild = node.lchild
		} else {
			node.parent.rchild = node.lchild
		}
		node.lchild.parent = node.parent
		node.lchild.color = "black"
		freeNode(node)
		return

	/*
		被删除节点只有一个子节点的时候，被删节点一定是黑，非nil子节点一定是红
		     X-> a●                             b●
		         / \             --------\      / \
		       nil  ○b           --------/    nil  nil
	*/
	case node.lchild == nil && node.rchild != nil:
		if node.parent == nil {
			node.rchild.parent = nil
			node.rchild.color = "black"
			rbt.root = node.rchild
			freeNode(node)
			return
		}
		if node.parent.lchild == node {
			node.parent.lchild = node.rchild
		} else {
			node.parent.rchild = node.rchild
		}
		node.rchild.parent = node.parent
		node.rchild.color = "black"
		freeNode(node)
		return

	/*
	stage1:删除节点，保持黑色平衡，临时引入两个颜色的节点
		（在node的lchild和rchild都不是nil的大前提下，它一定有后继）
		1.后继就是右子节点，进入stage2
		2.后继不是右子节点，进入stage2
	stage2:去除两个颜色的节点
		1.traceNode的兄弟节点是红色，调整；继续进行stage2的判断
		2.traceNode的兄弟节点是黑色，兄弟节点的右孩子黑，左孩子黑，调整；继续进行stage2的判断
		3.traceNode的兄弟节点是黑色，兄弟节点的右孩子黑，左孩子红，调整；进入4
		4.traceNode的兄弟节点是黑色，兄弟节点的右孩子红，调整；结束
	 */
	case node.lchild != nil && node.rchild != nil:
		traceNode := rbt.deleteStage1(node)
		if traceNode == nil {
			return
		}
		rbt.deleteStage2(traceNode)
	}
}

func (rbt *RBTree) deleteStage1(node *TreeNode) (traceNode *TreeNode) {
	defer func() {
		node.color = "no mean"
		node.lchild = nil
		node.rchild = nil
		node.parent = nil
	}()

	//（在node的lchild和rchild都不是nil的大前提下，它一定有后继）
	successorNode := rbt.GetSuccessor(node.data)
	oriColor := successorNode.color //保存后继节点原来的颜色
	successorNode.color = node.color //后继节点的颜色与node一样

	if node.rchild == successorNode {
		if node.parent == nil {
			successorNode.parent = nil
			rbt.root = successorNode
			rbt.root.color = "black"
		} else {
			successorNode.parent = node.parent
			if node == node.parent.lchild {
				node.parent.lchild = successorNode
			} else {
				node.parent.rchild = successorNode
			}
		}

		successorNode.lchild = node.lchild
		node.lchild.parent = successorNode

		if oriColor == "red" {
			return nil
		} else {
			//successorNode的右孩子，要么是红色非nil节点；要么是黑色nil节点
			if successorNode.rchild == nil {
				successorNode.rchild = &TreeNode{
					parent: successorNode,
					color: "black",
				}
			}
			successorNode.rchild.color += "+black"

			traceNode = successorNode.rchild
			return traceNode
		}
	} else {
		successorNodeParent := successorNode.parent
		successorNodeRChild := successorNode.rchild
		if node.parent == nil {
			successorNode.parent = nil
			rbt.root = successorNode
			rbt.root.color = "black"
		} else {
			successorNode.parent = node.parent
			if node == node.parent.lchild {
				node.parent.lchild = successorNode
			} else {
				node.parent.rchild = successorNode
			}
		}

		successorNode.lchild = node.lchild
		node.lchild.parent = successorNode

		node.rchild.parent = successorNode
		successorNode.rchild = node.rchild

		if oriColor == "red" {
			successorNodeParent.lchild = nil
			return nil
		} else {
			if successorNodeRChild == nil {
				successorNodeRChild = &TreeNode{
					color: "black",
				}
			}
			successorNodeRChild.parent = successorNodeParent
			successorNodeParent.lchild = successorNodeRChild
			successorNodeRChild.color += "+black"

			traceNode = successorNodeRChild
			return traceNode
		}
	}
}

func (rbt *RBTree) deleteStage2(traceNode *TreeNode) {
	var brother *TreeNode
	for traceNode != rbt.root && strings.Contains(traceNode.color, "+black") {
		if traceNode.parent.lchild == traceNode && traceNode.parent.rchild != nil {
			brother = traceNode.parent.rchild
			lCase1to4:
				for {
					switch {
					case brother.color == "red":
						brother.color = "black"
						traceNode.parent.color = "red"
						rbt.LeftRotate(traceNode.parent)
					case brother.color == "black" && brother.lchild != nil && brother.lchild.color == "black" && brother.rchild != nil && brother.rchild.color == "black":
						brother.color = "red"
						traceNode.color = strings.Split(traceNode.color, "+")[0]
						traceNode = traceNode.parent
						traceNode.color += "+black"
						if traceNode.color == "black" {
							traceNode.parent.lchild = nil
						}
						//traceNode变更，要重新找brother，所以break内层for循环
						break lCase1to4
					case brother.color == "black" && brother.lchild != nil && brother.lchild.color == "red" && brother.rchild != nil && brother.rchild.color == "black":
						brother.color = "red"
						brother.lchild.color = "black"
						rbt.RightRotate(brother)
						fallthrough
					case brother.color == "black" && brother.rchild != nil && brother.rchild.color == "red":
						brother.color = traceNode.parent.color
						brother.parent.color = "black"
						brother.rchild.color = "black"
						traceNode.color = strings.Split(traceNode.color, "+")[0]
						rbt.LeftRotate(brother.parent)
						if traceNode.color == "black" {
							traceNode.parent.lchild = nil
						}
						return
					}
				}
		} else if traceNode.parent.rchild == traceNode && traceNode.parent.lchild != nil {
			brother = traceNode.parent.lchild
			rCase1to4:
				for {
					switch {
					case brother.color == "red":
						brother.color = "black"
						traceNode.parent.color = "red"
						rbt.RightRotate(traceNode.parent)
					case brother.color == "black" && brother.lchild != nil && brother.lchild.color == "black" && brother.rchild != nil && brother.rchild.color == "black":
						brother.color = "red"
						traceNode.color = strings.Split(traceNode.color, "+")[0]
						traceNode = traceNode.parent
						traceNode.color += "+black"
						if traceNode.color == "black" {
							traceNode.parent.lchild = nil
						}
						//traceNode变更，要重新找brother，所以break内层for循环
						break rCase1to4
					case brother.color == "black" && brother.lchild != nil && brother.lchild.color == "black" && brother.rchild != nil && brother.rchild.color == "red":
						brother.color = "red"
						brother.rchild.color = "black"
						rbt.LeftRotate(brother)
						fallthrough
					case brother.color == "black" && brother.lchild != nil && brother.lchild.color == "red":
						brother.color = traceNode.parent.color
						brother.parent.color = "black"
						brother.lchild.color = "black"
						traceNode.color = strings.Split(traceNode.color, "+")[0]
						rbt.RightRotate(brother.parent)
						if traceNode.color == "black" {
							traceNode.parent.lchild = nil
						}
						return
					}
				}
		} else {
			return
		}
	}
}

func (rbt *RBTree) Delete(data float64) {
	var (
		deleteNode func(node *TreeNode)
		node       *TreeNode = rbt.Search(data)
		parent     *TreeNode
		revise     string
	)

	if node == nil {
		return
	} else {
		parent = node.parent
	}

	//下面这小块代码用来判断替代被删节点位置的节点是哪个后代
	if node.lchild == nil && node.rchild == nil {
		revise = "none"
	} else if parent == nil {
		revise = "root"
	} else if node == parent.lchild {
		revise = "left"
	} else if node == parent.rchild {
		revise = "right"
	}

	deleteNode = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.lchild == nil && node.rchild == nil {
			//如果要删除的节点没有孩子，直接删掉它就可以(毫无挂念~.~!)
			if node == rbt.root {
				rbt.root = nil
			} else {
				if node.parent.lchild == node {
					node.parent.lchild = nil
				} else {
					node.parent.rchild = nil
				}
			}

		} else if node.lchild != nil && node.rchild == nil {
			//如果要删除的节点只有左孩子或右孩子，让这个节点的父节点指向它的指针指向它的
			//孩子即可
			if node == rbt.root {
				node.lchild.parent = nil
				rbt.root = node.lchild
			} else {
				node.lchild.parent = node.parent
				if node.parent.lchild == node {
					node.parent.lchild = node.lchild
				} else {
					node.parent.rchild = node.lchild
				}
			}

		} else if node.lchild == nil && node.rchild != nil {
			if node == rbt.root {
				node.rchild.parent = nil
				rbt.root = node.rchild
			} else {
				node.rchild.parent = node.parent
				if node.parent.lchild == node {
					node.parent.lchild = node.rchild
				} else {
					node.parent.rchild = node.rchild
				}
			}

		} else {
			//如果要删除的节点既有左孩子又有右孩子，就把这个节点的直接后继的值赋给这个节
			//点，然后删除直接后继节点即可
			successor := rbt.GetSuccessor(node.data)
			node.data = successor.data
			node.color = successor.color
			deleteNode(successor)
		}
	}

	deleteNode(node)
	if node.color == "black" {
		if revise == "root" {
			rbt.deleteBalanceFixup(rbt.root)
		} else if revise == "left" {
			rbt.deleteBalanceFixup(parent.lchild)
		} else if revise == "right" {
			rbt.deleteBalanceFixup(parent.rchild)
		}
	}
	//至于为什么删除的为红节点时不用调节平衡，那本黑色的《算法导论(第二版)》是这么解释的：
	//当删除红节点时
	//(1) 树中个节点的黑高度都没有变化
	//(2) 不存在两个相邻的红色节点
	//(3) 如果删除的节点是红的，就不可能是根，所以根依然是黑色的
}

//这个函数用于在红黑树执行插入操作后，修复红黑性质
func (rbt *RBTree) insertBalanceFixup(insertnode *TreeNode) {
	var uncle *TreeNode

	for insertnode.color == "red" && insertnode.parent.color == "red" {
		//获取新插入的节点的叔叔节点(与父节点同根的另一个节点)
		if insertnode.parent == insertnode.parent.parent.lchild {
			uncle = insertnode.parent.parent.rchild
		} else {
			uncle = insertnode.parent.parent.lchild
		}

		if uncle != nil && uncle.color == "red" {
			//如果叔叔节点是红色，就按照如下图所示变化(●->黑, ○->红)：
			//     |                        |
			//    1●                       1○ <-new node ptr come here
			//    / \       --------\      / \
			//  2○   ○3     --------/    2●   ●3
			//  /                        /
			//4○ <-new node ptr        4○
			//
			//这种情况可以一直循环，知道new node ptr指到root时退出(root颜色依然为黑)
			uncle.color, insertnode.parent.color = "black", "black"
			insertnode = insertnode.parent.parent
			if insertnode == rbt.root || insertnode == nil {
				return
			}
			insertnode.color = "red"

		} else {
			//1.插入节点的父节点 是 插入节点的祖父节点 的左孩子（说明要绕祖父节点右旋）
			//	1）插入节点 是 插入节点的父节点 的左孩子
			//	2）插入节点 是 插入节点的父节点 的右孩子（多一步 左旋。因为马上要右旋，新插入节点最好位于左子树。）
			//2.插入节点的父节点 是 插入节点的祖父节点 的右孩子（说明要绕祖父节点左旋）
			//	1）插入节点 是 插入节点的父节点 的左孩子（多一步 右旋。因为马上要左旋，新插入节点最好位于右子树。）
			//	2）插入节点 是 插入节点的父节点 的右孩子
			//《如果叔叔节点为空或叔叔节点为黑色，就按照如下图所示变化：
			//     |                        |
			//    1● <-right rotate        2●
			//    / \       --------\      / \
			//  2○   ●3     --------/    4○   ○1
			//  /                              \
			//4○ <-new node ptr                 ●3
			//
			//当然，这只是叔叔节点为黑时的一种情况，如果上图的节点4为2节点的右孩子，则
			//可以先在2节点处向左旋转，这样就转换成了上图这种情况了。另外两种情况自己想
			//想就明白了》
			if insertnode.parent == insertnode.parent.parent.lchild {
				if insertnode == insertnode.parent.rchild {
					insertnode = insertnode.parent
					rbt.LeftRotate(insertnode)
				}
				insertnode = insertnode.parent
				insertnode.color = "black"
				insertnode = insertnode.parent
				insertnode.color = "red"
				rbt.RightRotate(insertnode)

			} else {
				//     |                                  |
				//    1● <-left rotate                  2●
				//    / \       ------------------\      / \
				//  3●  ○2     ------------------/   1○  ○4
				//        \                           /
				//        ○4 <-new node ptr        3●
				//
				if insertnode == insertnode.parent.lchild {
					insertnode = insertnode.parent
					rbt.RightRotate(insertnode)
				}
				insertnode = insertnode.parent
				insertnode.color = "black"
				insertnode = insertnode.parent
				insertnode.color = "red"
				rbt.LeftRotate(insertnode)
			}
			return
		}
	}
}

//这个函数用于在红黑树执行删除操作后，修复红黑性质
func (rbt *RBTree) deleteBalanceFixup(node *TreeNode) {
	var brother *TreeNode

	for node != rbt.root && node.color == "black" {
		//删除时的红黑修复需要考虑四种情况(下面的“X”指取代了被删节点位置的新节点)
		// (1) X的兄弟节点是红色的：
		//        |                              |
		//       1●                             3●
		//       / \             --------\      / \
		// X-> 2●   ○3 <-brother --------/    1○   ●5
		//         / \                        / \
		//       4●   ●5                X-> 2●   ●4
		//
		// (2) X的兄弟节点是黑色的，而且兄弟节点的两个孩子都是黑色的：
		//        |                              |
		//       1○                         X-> 1○
		//       / \             --------\      / \
		// X-> 2●   ●3 <-brother --------/    2●   ○3
		//         / \                            / \
		//       4●   ●5                        4●   ●5
		//
		// (3) X的兄弟节点是黑色的，兄弟的左孩子是红色的，右孩子是黑色的：
		//        |                              |
		//       1○                             1○
		//       / \             --------\      / \
		// X-> 2●   ●3 <-brother --------/ X->2●   ●4
		//         / \                              \
		//       4○   ●5                             ○3
		//                                            \
		//                                             ●5
		//
		// (4) X的兄弟节点是黑色的，兄弟的右孩子是红色的：
		//        |                              |
		//       1○                             3○   X->root and loop while end
		//       / \             --------\      / \
		// X-> 2●   ●3 <-brother --------/    1●   ●5
		//         / \                        / \
		//       4○   ○5                    2●   ○4
		//
		//以上是兄弟节点在X右边时的情况，在X左边是取相反即可!
		if node.parent.lchild == node && node.parent.rchild != nil {
			brother = node.parent.rchild
			if brother.color == "red" {
				brother.color = "black"
				node.parent.color = "red"
				rbt.LeftRotate(node.parent)
			} else if brother.color == "black" && brother.lchild != nil && brother.lchild.color == "black" && brother.rchild != nil && brother.rchild.color == "black" {
				brother.color = "red"
				node = node.parent
			} else if brother.color == "black" && brother.lchild != nil && brother.lchild.color == "red" && brother.rchild != nil && brother.rchild.color == "black" {
				brother.color = "red"
				brother.lchild.color = "black"
				rbt.RightRotate(brother)
			} else if brother.color == "black" && brother.rchild != nil && brother.rchild.color == "red" {
				brother.color = "red"
				brother.rchild.color = "black"
				brother.parent.color = "black"
				rbt.LeftRotate(brother.parent)
				node = rbt.root
			}

		} else if node.parent.rchild == node && node.parent.lchild != nil {
			brother = node.parent.lchild
			if brother.color == "red" {
				brother.color = "black"
				node.parent.color = "red"
				rbt.RightRotate(node.parent)
			} else if brother.color == "black" && brother.lchild != nil && brother.lchild.color == "black" && brother.rchild != nil && brother.rchild.color == "black" {
				brother.color = "red"
				node = node.parent
			} else if brother.color == "black" && brother.lchild != nil && brother.lchild.color == "black" && brother.rchild != nil && brother.rchild.color == "red" {
				brother.color = "red"
				brother.rchild.color = "black"
				rbt.LeftRotate(brother)
			} else if brother.color == "black" && brother.lchild != nil && brother.lchild.color == "red" {
				brother.color = "red"
				brother.lchild.color = "black"
				brother.parent.color = "black"
				rbt.RightRotate(brother.parent)
				node = rbt.root
			}
		} else {
			return
		}
	}
}

func (rbt RBTree) GetRoot() *TreeNode {
	if rbt.root != nil {
		return rbt.root
	}
	return nil
}

func (rbt RBTree) IsEmpty() bool {
	if rbt.root == nil {
		return true
	}
	return false
}

func (rbt RBTree) InOrderTravel(op common.Operate) {
	var inOrderTravel func(node *TreeNode)

	inOrderTravel = func(node *TreeNode) {
		if node != nil {
			inOrderTravel(node.lchild)
			op(fmt.Sprintf("%f, %s", node.data, node.color))
			//fmt.Printf("%g ", node.data)
			inOrderTravel(node.rchild)
		}
	}

	inOrderTravel(rbt.root)
}

//寻找指定节点
func (rbt RBTree) Search(data float64) *TreeNode {
	//和Add操作类似，只要按照比当前节点小就往左孩子上拐，比当前节点大就往右孩子上拐的思路
	//一路找下去，知道找到要查找的值返回即可
	rbt.cur = rbt.root
	for {
		if rbt.cur == nil {
			return nil
		}

		if data < rbt.cur.data {
			rbt.cur = rbt.cur.lchild
		} else if data > rbt.cur.data {
			rbt.cur = rbt.cur.rchild
		} else {
			return rbt.cur
		}
	}
}

func (rbt RBTree) GetDeepth() int {
	var getDeepth func(node *TreeNode) int

	getDeepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.lchild == nil && node.rchild == nil {
			return 1
		}
		var (
			ldeepth int = getDeepth(node.lchild)
			rdeepth int = getDeepth(node.rchild)
		)
		if ldeepth > rdeepth {
			return ldeepth + 1
		} else {
			return rdeepth + 1
		}
	}

	return getDeepth(rbt.root)
}

func (rbt RBTree) GetMin() float64 {
	//根据二叉查找树的性质，树中最左边的节点就是值最小的节点
	if rbt.root == nil {
		return -1
	}
	rbt.cur = rbt.root
	for {
		if rbt.cur.lchild != nil {
			rbt.cur = rbt.cur.lchild
		} else {
			return rbt.cur.data
		}
	}
}

func (rbt RBTree) GetMax() float64 {
	//根据二叉查找树的性质，树中最右边的节点就是值最大的节点
	if rbt.root == nil {
		return -1
	}
	rbt.cur = rbt.root
	for {
		if rbt.cur.rchild != nil {
			rbt.cur = rbt.cur.rchild
		} else {
			return rbt.cur.data
		}
	}
}

func (rbt RBTree) GetPredecessor(data float64) *TreeNode {
	getMax := func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		for {
			if node.rchild != nil {
				node = node.rchild
			} else {
				return node
			}
		}
	}

	node := rbt.Search(data)
	if node != nil {
		if node.lchild != nil {
			//如果这个节点有左孩子，那么它的直接前驱就是它左子树的最右边的节点，因为比这
			//个节点值小的节点都在左子树，而这些节点中值最大的就是这个最右边的节点
			return getMax(node.lchild)
		} else {
			//如果这个节点没有左孩子，那么就沿着它的父节点找，知道某个父节点的父节点的右
			//孩子就是这个父节点，那么这个父节点的父节点就是直接前驱
			for {
				if node == nil || node.parent == nil {
					break
				}
				if node == node.parent.rchild {
					return node.parent
				}
				node = node.parent
			}
		}
	}

	return nil
}

//寻找节点的后继节点
//只要根据data能找到节点，那么返回nil就只有一种情况：该节点是序列中的最大节点
func (rbt RBTree) GetSuccessor(data float64) *TreeNode {
	getMin := func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		for {
			if node.lchild != nil {
				node = node.lchild
			} else {
				return node
			}
		}
	}

	//参照寻找直接前驱的函数对比着看
	node := rbt.Search(data)
	if node != nil {
		if node.rchild != nil {
			return getMin(node.rchild)
		} else {
			for {
				if node == nil || node.parent == nil {
					break
				}
				if node == node.parent.lchild {
					return node.parent
				}
				node = node.parent
			}
		}
	}

	return nil
}

func (rbt *RBTree) Clear() {
	rbt.root = nil
	rbt.cur = nil
	rbt.create = nil
}

/**
 * 旋转图解(以向左旋转为例)：
 *     |                               |
 *     ○ <-left rotate                 ●
 *      \              ----------\    / \
 *       ●             ----------/   ○   ●r
 *      / \                           \
 *    l●   ●r                         l●
 *
 *
 *
 *     |                               |
 *     ○ <-left rotate                 ●
 *      \              ----------\    / \
 *       ●             ----------/   ○   ●
 *        \                           \
 *         ●                          nil <-don't forget it should be nil
 */
func (rbt *RBTree) LeftRotate(node *TreeNode) {
	if node.rchild == nil {
		return
	}

	//1.移除node.rChild.lChild。因为左旋（逆时针旋）后，lChild与node会“干涉”。
	//后面马上修改node.rChild，所以在修改之前，先把它保存
	rightChild := node.rchild
	//《将要旋转的节点的右孩子的左孩子赋给这个节点的右孩子，这里最好按如下3行代码的顺序写，》
	//《否则该节点的右孩子的左孩子为nil时，很容易忘记把这个节点的右孩子也置为nil.》
	node.rchild = rightChild.lchild
	if node.rchild != nil {
		node.rchild.parent = node
	}

	//2.处理node.parent。因为左旋（逆时针旋）后，node.parent.lChild/node.parent.rChild要发生变化。
	//《让要旋转的节点的右孩子的父节点指针指向当前节点父节点。如果父节点是根节点要特别处理》
	rightChild.parent = node.parent
	if node.parent == nil {
		rbt.root = rightChild
	} else {
		if node.parent.lchild == node {
			node.parent.lchild = rightChild
		} else {
			node.parent.rchild = rightChild
		}
	}

	//3.添加node.rChild.lChild。因为左旋（逆时针旋）后，node与node.rChild的父子关系，会变化。
	//《上面的准备工作完毕了，就可以开始旋转了，让要旋转的节点的右孩子的左孩子指向该节点，》
	//《别忘了把这个被旋转的节点的父节点指针指向新的父节点》
	rightChild.lchild = node
	node.parent = rightChild
}

func (rbt *RBTree) RightRotate(node *TreeNode) {
	//《向右旋转的过程与LeftRotate()正好相反》
	if node.lchild == nil {
		return
	}

	//1.移除旋转后的“干涉”。
	leftChild := node.lchild
	node.lchild = leftChild.rchild
	if node.lchild != nil {
		node.lchild.parent = node
	}

	//2.处理node.parent。
	leftChild.parent = node.parent
	if node.parent == nil {
		rbt.root = leftChild
	} else {
		if node.parent.lchild == node {
			node.parent.lchild = leftChild
		} else {
			node.parent.rchild = leftChild
		}
	}

	//3.之前“干涉”位置被移除，现在添加，也就是真正的旋转操作。
	leftChild.rchild = node
	node.parent = leftChild
}
