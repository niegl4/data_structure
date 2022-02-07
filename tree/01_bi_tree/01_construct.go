package _1_bi_tree

import "errors"

/*
七
已知二叉树的前序，中序遍历结果，重建二叉树。二叉树中不包含重复数字。
todo: 已知中序，后续，重建二叉树。
*/
func construct(preOrder, inOrder []int) (*BiTree, error) {
	preLen := len(preOrder)
	inLen := len(inOrder)
	if preLen == 0 || inLen == 0 || preLen != inLen {
		return nil, errors.New("input invalid")
	}
	return constructCore(preOrder, 0, preLen-1, inOrder, 0, inLen-1)
}

func constructCore(preOrder []int, preStart, preEnd int, inOrder []int, inStart, inEnd int) (*BiTree, error) {
	node := &BiTree{
		Data: preOrder[preStart],
	}
	if preEnd == preStart && inEnd == inStart {
		if preOrder[preStart] == inOrder[inStart] {
			return node, nil
		} else {
			return nil, errors.New("input invalid")
		}
	}

	inOrderIndex := inStart
	for inOrderIndex <= inEnd {
		if inOrder[inOrderIndex] == preOrder[preStart] {
			break
		} else {
			inOrderIndex++
		}
	}
	if inOrderIndex > inEnd {
		return nil, errors.New("input invalid")
	}
	leftLen := inOrderIndex - inStart
	rightLen := inEnd - inOrderIndex
	if leftLen > 0 {
		left, err := constructCore(preOrder, preStart+1, preStart+leftLen,
			inOrder, inStart, inOrderIndex-1)
		if err != nil {
			return nil, err
		} else {
			node.lChild = left
		}
	}
	if rightLen > 0 {
		right, err := constructCore(preOrder, preStart+leftLen+1, preEnd,
			inOrder, inOrderIndex+1, inEnd)
		if err != nil {
			return nil, err
		} else {
			node.rChild = right
		}
	}
	return node, nil
}
