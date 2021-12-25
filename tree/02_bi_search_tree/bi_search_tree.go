package _2_bi_search_tree

type TreeNode struct {
	data   float64
	lChild *TreeNode
	rChild *TreeNode
	parent *TreeNode
}

type BiSTree struct {
	root   *TreeNode
	cur    *TreeNode
	create *TreeNode
}

func (bit *BiSTree) Add(data float64) {
	bit.create = &TreeNode{
		data: data,
	}

	if !bit.IsEmpty() {
		bit.cur = bit.root
		for {
			if data < bit.cur.data {
				if bit.cur.lChild == nil {
					bit.cur.lChild = bit.create
					bit.create.parent = bit.cur
					break
				} else {
					bit.cur = bit.cur.lChild
				}
			} else if data > bit.cur.data {
				if bit.cur.rChild == nil {
					bit.cur.rChild = bit.create
					bit.create.parent = bit.cur
					break
				} else {
					bit.cur = bit.cur.rChild
				}
			} else {
				return
			}
		}
	} else {
		bit.root = bit.create
		return
	}
}

func (bit *BiSTree) Delete(data float64) {
	node := bit.Search(data)
	if node == nil {
		return
	}

	switch {
	case node.lChild == nil && node.rChild == nil:
		if node.parent == nil {
			bit.root = nil
			return
		}
		if node == node.parent.lChild {
			node.parent.lChild = nil
		} else {
			node.parent.rChild = nil
		}

	case node.lChild != nil && node.rChild == nil:
		if node.parent == nil {
			bit.root = node.lChild
			bit.root.parent = nil
			return
		}
		if node == node.parent.lChild {
			node.parent.lChild = node.lChild
		} else {
			node.parent.rChild = node.lChild
		}
		node.lChild.parent = node.parent

	case node.lChild == nil && node.rChild != nil:
		if node.parent == nil {
			bit.root = node.rChild
			bit.root.parent = nil
			return
		}
		if node == node.parent.lChild {
			node.parent.lChild = node.rChild
		} else {
			node.parent.rChild = node.rChild
		}
		node.rChild.parent = node.parent

	case node.lChild != nil && node.rChild != nil:
		successorNode := bit.GetSuccessor(data)
		if node.rChild == successorNode { //后继节点就是node.rChild
			if node.parent == nil {
				successorNode.parent = nil
				bit.root = successorNode
			} else {
				if node == node.parent.lChild {
					node.parent.lChild = successorNode
				} else {
					node.parent.rChild = successorNode
				}
				successorNode.parent = node.parent
			}

			successorNode.lChild = node.lChild
			node.lChild.parent = successorNode

		} else { //后继节点不是node.rChild
			successorNodeParent := successorNode.parent
			successorNodeRChild := successorNode.rChild

			if node.parent == nil {
				successorNode.parent = nil
				bit.root = successorNode
			} else {
				if node == node.parent.lChild {
					node.parent.lChild = successorNode
				} else {
					node.parent.rChild = successorNode
				}
				successorNode.parent = node.parent
			}

			successorNode.lChild = node.lChild
			node.lChild.parent = successorNode

			successorNode.rChild = node.rChild
			node.rChild.parent = successorNode

			successorNodeParent.lChild = successorNodeRChild
			if successorNodeRChild != nil { //后继节点的右孩子可能是nil
				successorNodeRChild.parent = successorNodeParent
			}
		}
	}
}

func (bit *BiSTree) Search(data float64) *TreeNode {
	bit.cur = bit.root
	for {
		if bit.cur == nil {
			return nil
		}

		if data < bit.cur.data {
			bit.cur = bit.cur.lChild
		} else if data > bit.cur.data {
			bit.cur = bit.cur.rChild
		} else {
			return bit.cur
		}
	}
}

func (bit *BiSTree) GetSuccessor(data float64) *TreeNode {
	getMin := func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		for {
			if node.lChild == nil {
				return node
			} else {
				node = node.lChild
			}
		}
	}

	node := bit.Search(data)
	if node == nil {
		return nil
	}
	if node.rChild != nil {
		return getMin(node.rChild)
	} else {
		for {
			if node == nil || node.parent == nil {
				return nil
			}
			if node == node.parent.lChild {
				return node.parent
			} else {
				node = node.parent
			}
		}
	}
}

func (bit *BiSTree) InOrderTravel() (res []interface{}) {
	var inOrderTravel func(node *TreeNode) (res []interface{})
	inOrderTravel = func(node *TreeNode) (res []interface{}) {
		if node == nil {
			return res
		}
		res = append(res, inOrderTravel(node.lChild)...)
		res = append(res, node.data)
		res = append(res, inOrderTravel(node.rChild))
		return res
	}

	return inOrderTravel(bit.root)
}

func (bit BiSTree) IsEmpty() bool {
	if bit.root == nil {
		return true
	}
	return false
}
