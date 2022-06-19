package dynamic_algorithm

/*
动态规划比较适合用来求解最优问题，比如求最大值、最小值等等。它可以非常显著地降低时间复杂度，提高代码的执行效率。

我们把问题分解为多个阶段，每个阶段对应一个决策。
我们记录每一个阶段可达的状态集合（去掉重复的），然后通过当前阶段的状态集合，来推导下一个阶段的状态集合，动态地往前推进。
维护状态集合最大的好处就是：避免了每层状态个数的指数级增长。
*/

/*
01
一组物品，装入包中。在满足背包最大重量限制的前提下，背包中物品总重是多少？
*/

//回溯算法实现，时间复杂O(2^n)
func packageByBacktrack(objs []int, weightThreshold int) int {
	maxW := 0
	packageByBacktrackCore(objs, 0, len(objs),
		0, weightThreshold, &maxW)
	return maxW
}

func packageByBacktrackCore(objs []int, idx, n, cw, weightThreshold int,
	maxW *int) {
	if cw == weightThreshold || idx == n {
		if cw > *maxW {
			*maxW = cw
		}
		return
	}

	packageByBacktrackCore(objs, idx+1, n, cw, weightThreshold, maxW)
	if cw+objs[idx] <= weightThreshold {
		packageByBacktrackCore(objs, idx+1, n, cw+objs[idx], weightThreshold, maxW)
	}
}

//动态规划实现，【二维数组】实现状态信息表，时间复杂度O(物品个数*(重量阈值+1))，空间复杂度O(物品个数*(重量阈值+1))
func packageWithMatrix(objs []int, weightThreshold int) int {
	if weightThreshold < 0 {
		return -1
	}
	if len(objs) < 1 {
		return 0
	}

	matrix := make([][]bool, len(objs)) //二维状态信息表
	for i := 0; i < len(objs); i++ {
		matrix[i] = make([]bool, weightThreshold+1)
	}

	//第一行填写
	matrix[0][0] = true
	if objs[0] <= weightThreshold {
		matrix[0][objs[0]] = true
	}

	//从第二行开始，它要参考上一行的状态
	for i := 1; i < len(objs); i++ {
		for j := 0; j <= weightThreshold; j++ { //继承上一行的状态
			if matrix[i-1][j] {
				matrix[i][j] = true
			}
		}

		for j := 0; j <= weightThreshold-objs[i]; j++ { //在不超重前提下，上一行装了的，生成本行状态
			if matrix[i-1][j] {
				matrix[i][j+objs[i]] = true
			}
		}
	}

	//最大可能质量，保存在最后一行的最后一个为true的位置
	for j := weightThreshold; j >= 0; j-- {
		if matrix[len(objs)-1][j] {
			return j
		}
	}
	return 0
}

//动态规划实现，【一维数组】实现状态信息表，时间复杂度O(物品个数*(重量阈值+1))，空间复杂度O(重量阈值+1)
func packageWithArray(objs []int, weightThreshold int) int {
	if weightThreshold < 0 {
		return -1
	}
	if len(objs) < 1 {
		return 0
	}

	arr := make([]bool, weightThreshold+1)

	arr[0] = true
	if objs[0] <= weightThreshold {
		arr[objs[0]] = true
	}

	//使用matrix时，需要继承上一行的状态。使用array时，不用做处理，直接就继承。
	for i := 1; i < len(objs); i++ {
		//最精髓的地方：从后往前修改状态位。因为是array，读本行写本行，而不是matrix，读上行写本行。
		for j := weightThreshold - objs[i]; j >= 0; j-- {
			if arr[j] {
				arr[j+objs[i]] = true
			}
		}
	}

	for j := weightThreshold; j >= 0; j-- {
		if arr[j] {
			return j
		}
	}
	return 0
}

/*
02
一组物品，装入包中。在满足背包最大重量限制的前提下，背包中物品【总价值】是多少？
*/
func packageValueWithArray(objs, values []int, weightThreshold int) int {
	if len(objs) != len(values) {
		return -1
	}
	if weightThreshold < 0 {
		return -1
	}
	if len(objs) < 1 {
		return 0
	}

	arr := make([]int, weightThreshold+1)
	for i := 0; i <= weightThreshold; i++ {
		arr[i] = -1
	}

	arr[0] = 0
	if objs[0] <= weightThreshold {
		arr[objs[0]] = values[0]
	}

	//使用matrix时，需要继承上一行的状态。使用array时，不用做处理，直接就继承。
	for i := 1; i < len(objs); i++ {
		//最精髓的地方：从后往前修改状态位。因为是array，读本行写本行，而不是matrix，读上行写本行。
		for j := weightThreshold - objs[i]; j >= 0; j-- {
			if arr[j] >= 0 {
				tmpVal := arr[j] + values[i]
				if tmpVal > arr[j+objs[i]] {
					arr[j+objs[i]] = tmpVal
				}
			}
		}
	}

	res := 0
	for j := 0; j <= weightThreshold; j++ {
		if arr[j] >= 0 {
			res = arr[j]
		}
	}
	return res
}
