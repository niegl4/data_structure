package _1_bi_tree

import "errors"

/*
七
已知二叉树的前序，中序遍历结果，重建二叉树。二叉树中不包含重复数字。
*/
func constructByPreAndIn(preOrder, inOrder []int) (*BiTree, error) {
	preLen := len(preOrder)
	inLen := len(inOrder)
	if preLen == 0 || inLen == 0 || preLen != inLen {
		return nil, errors.New("input invalid")
	}
	return constructCoreV1(preOrder, 0, preLen-1, inOrder, 0, inLen-1)
}

func constructCoreV1(preOrder []int, preStart, preEnd int, inOrder []int, inStart, inEnd int) (*BiTree, error) {
	node := &BiTree{
		Data: preOrder[preStart],
	}
	//递归终止条件：两个区间都只有一个元素，并且值相等
	if preEnd == preStart && inEnd == inStart {
		if preOrder[preStart] == inOrder[inStart] {
			return node, nil
		} else {
			return nil, errors.New("input invalid")
		}
	}

	//探索中序遍历的“中点”，比如：第一次遍历的中点就是root节点。
	inOrderIndex := inStart
	for inOrderIndex <= inEnd {
		if inOrder[inOrderIndex] == preOrder[preStart] {
			break
		}
		inOrderIndex++
	}
	if inOrderIndex > inEnd {
		return nil, errors.New("input invalid")
	}

	//前序，中序区间，分别划分左右区间
	leftLen := inOrderIndex - inStart
	rightLen := inEnd - inOrderIndex
	if leftLen > 0 {
		left, err := constructCoreV1(preOrder, preStart+1, preStart+leftLen,
			inOrder, inStart, inOrderIndex-1)
		if err != nil {
			return nil, err
		} else {
			node.lChild = left
		}
	}
	if rightLen > 0 {
		right, err := constructCoreV1(preOrder, preStart+leftLen+1, preEnd,
			inOrder, inOrderIndex+1, inEnd)
		if err != nil {
			return nil, err
		} else {
			node.rChild = right
		}
	}
	return node, nil
}

//已知中序，后序，重建二叉树
func constructByPostAndIn(postOrder, inOrder []int) (*BiTree, error) {
	postLen := len(postOrder)
	inLen := len(inOrder)
	if postLen == 0 || inLen == 0 || postLen != inLen {
		return nil, errors.New("input invalid")
	}
	return constructCoreV2(postOrder, 0, postLen-1, inOrder, 0, inLen-1)
}

func constructCoreV2(postOrder []int, postStart, postEnd int, inOrder []int, inStart, inEnd int) (*BiTree, error) {
	node := &BiTree{
		Data: postOrder[postEnd],
	}

	if postStart == postEnd && inStart == inEnd {
		if postOrder[postStart] == inOrder[inStart] {
			return node, nil
		} else {
			return nil, errors.New("invalid input")
		}
	}

	inOrderIndex := inStart
	for inOrderIndex < inEnd {
		if inOrder[inOrderIndex] == postOrder[postEnd] {
			break
		}
		inOrderIndex++
	}

	leftLen := inOrderIndex - inStart
	rightLen := inEnd - inOrderIndex
	if leftLen > 0 {
		lChild, err := constructCoreV2(postOrder, postStart, postStart+leftLen-1,
			inOrder, inStart, inOrderIndex-1)
		if err != nil {
			return nil, err
		} else {
			node.lChild = lChild
		}
	}
	if rightLen > 0 {
		rChild, err := constructCoreV2(postOrder, postStart+leftLen, postEnd-1,
			inOrder, inOrderIndex+1, inEnd)
		if err != nil {
			return nil, err
		} else {
			node.rChild = rChild
		}
	}
	return node, nil
}
