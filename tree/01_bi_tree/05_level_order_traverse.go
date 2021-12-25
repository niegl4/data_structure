package _1_bi_tree

/*
三十二-1
不分行，层序打印二叉树，每一层从左到右
*/

// levelOrderTraverse1 二叉树的层序遍历，不分行
func (bi *BiTree) levelOrderTraverse1() (res []interface{}) {
	if bi == nil {
		return
	}
	res = make([]interface{}, 0)
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
		res = append(res, level...)
		//queue截取掉本层
		queue = queue[n:]
	}
	return res
}

/*
三十二-2
分行，层序打印二叉树，每一层从左到右
*/

// levelOrderTraverse2 二叉树的层序遍历，分行
func (bi *BiTree) levelOrderTraverse2() (res [][]interface{}) {
	if bi == nil {
		return
	}
	res = make([][]interface{}, 0)
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
		res = append(res, level)
		//queue截取掉本层
		queue = queue[n:]
	}
	return res
}

/*
三十二-3
"之"字打印二叉树：
	  ->
	 <---
	------>
   <--------

两个栈保存接下来的节点。层数从1开始计算，奇数层：左子节点先入栈；偶数层：右子节点先入栈。
*/

// levelOrderTraverse3 二叉树的"之"字层序遍历，分行
func (bi *BiTree) levelOrderTraverse3() (res [][]interface{}) {
	if bi == nil {
		return
	}
	res = make([][]interface{}, 0)
	levelResTmp := make([]interface{}, 0)

	oddStack := []*BiTree{bi}
	evenStack := make([]*BiTree, 0)

	level := 1
	for len(oddStack) > 0 || len(evenStack) > 0 {
		if level&0b1 == 1 {
			//出栈
			node := oddStack[len(oddStack)-1]
			oddStack = oddStack[:len(oddStack)-1]

			levelResTmp = append(levelResTmp, node.Data)
			//子节点入偶数栈,先左
			if node.lChild != nil {
				evenStack = append(evenStack, node.lChild)
			}
			if node.rChild != nil {
				evenStack = append(evenStack, node.rChild)
			}
			//奇数栈为空，层数加1，层结果合并入最终结果
			if len(oddStack) == 0 {
				var levelRes []interface{}
				levelRes = append(levelRes, levelResTmp...)
				res = append(res, levelRes)
				levelResTmp = levelResTmp[0:0]
				level++
			}
		} else {
			node := evenStack[len(evenStack)-1]
			evenStack = evenStack[:len(evenStack)-1]

			levelResTmp = append(levelResTmp, node.Data)
			//子节点入奇数栈，先右
			if node.rChild != nil {
				oddStack = append(oddStack, node.rChild)
			}
			if node.lChild != nil {
				oddStack = append(oddStack, node.lChild)
			}
			if len(evenStack) == 0 {
				var levelRes []interface{}
				levelRes = append(levelRes, levelResTmp...)
				res = append(res, levelRes)
				levelResTmp = levelResTmp[0:0]
				level++
			}
		}
	}
	return res
}
