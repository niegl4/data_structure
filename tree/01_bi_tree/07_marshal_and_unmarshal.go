package _1_bi_tree

/*
三十七
序列化二叉树，反序列化二叉树。

与第七题相比，二叉树中有值相等的节点，这是最大不同。
*/
func marshalBT(root *BiTree) (res []interface{}) {
	if root == nil {
		res = append(res, "nil")
		return
	}
	res = append(res, root.Data)
	res = append(res, marshalBT(root.lChild)...)
	res = append(res, marshalBT(root.rChild)...)
	return res
}

func unmarshalBT(sequence []interface{}) (root *BiTree) {
	if len(sequence) < 1 {
		return nil
	}
	root, _ = unmarshalBTCore(sequence, 0)
	return root
}

func unmarshalBTCore(sequence []interface{}, curIdx int) (*BiTree, int) {
	if curIdx >= len(sequence) || sequence[curIdx] == "nil" {
		curIdx++
		return nil, curIdx
	}
	node := &BiTree{
		Data:   sequence[curIdx],
		lChild: nil,
		rChild: nil,
	}
	curIdx++
	lChild, curIdx := unmarshalBTCore(sequence, curIdx)
	rChild, curIdx := unmarshalBTCore(sequence, curIdx)
	node.lChild = lChild
	node.rChild = rChild
	return node, curIdx
}
