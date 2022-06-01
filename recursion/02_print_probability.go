package recursion

import (
	"fmt"
	"math"
)

/*
六十
n个骰子，朝上一面的和为s，打印s的所有可能的值出现的概率。
*/

var maxVal = 6

/*
v1是递归实现。
1.maxVal全局变量，应对多面骰子，可以是6个面，也可以是5面，等等。
2.巧用数组为map，key为sum，value为出现的次数。而且还有偏移量。
3.临时变量sum，通过函数形参，层层传递。

*
递归核心函数：第n个筛子，出现的点数，当前的概率数组。外加一个操作概率数组需要的偏移量。
*/
func printProbabilityV1(n int) {
	if n < 1 {
		return
	}

	maxSum := n * maxVal
	//sum共有(n*maxVal - n + 1)种可能,以数组为map,key为sum,value为出现的次数。
	//另外sum不可能小于n，所以读写数组时，再偏移n
	probabilities := make([]int, maxSum-n+1)

	probability(n, probabilities)

	total := math.Pow(float64(maxVal), float64(n))
	for i := n; i <= maxSum; i++ {
		fmt.Printf("%d, %f\n", i, float64(probabilities[i-n])/total)
	}
}

//第n个骰子的值，分别从1到maxVal
func probability(n int, probabilities []int) {
	for i := 1; i <= maxVal; i++ {
		probabilityCore(n, n, i, probabilities)
	}
}

//original:几个骰子。必传，因为读写数组需要该偏移量
//current:当前是第几个骰子
//sum:目前为止，已经统计的骰子的点数和
func probabilityCore(original, current, sum int, probabilities []int) {
	if current == 1 {
		probabilities[sum-original]++
	} else {
		for i := 1; i <= maxVal; i++ {
			//下一个骰子，分别从1到maxVal，sum同步累加
			probabilityCore(original, current-1, sum+i, probabilities)
		}
	}
}

/*
v2是循环实现。
1.巧用数组为map，key为sum，value为出现的次数。没有偏移量。
2.两个数组交替使用。一个数组计算结果，是另一个数组的计算基础，避免了重复计算。
仍然是临时变量保存计算结果，循环代替递归。

***
*/
func printProbabilityV2(n int) {
	if n < 1 {
		return
	}

	probabilities := make([][]int, 2)
	probabilities[0] = make([]int, maxVal*n+1)
	probabilities[1] = make([]int, maxVal*n+1)
	flag := 0

	//第1个骰子
	for i := 1; i <= maxVal; i++ {
		probabilities[flag][i] = 1
	}

	for k := 2; k <= n; k++ {
		//新数组中，不可能出现的sum置为0
		for i := 0; i < k; i++ {
			probabilities[1-flag][i] = 0
		}

		//可能出现的sum从k到maxVal*k
		for i := k; i <= maxVal*k; i++ {
			//旧数据写0，为更新做准备
			probabilities[1-flag][i] = 0

			//第k个骰子，从1变化到maxVal。它的值必须比当前sum（即，i）小。
			for j := 1; j <= maxVal && j < i; j++ {
				probabilities[1-flag][i] += probabilities[flag][i-j]
			}
		}

		flag = 1 - flag
	}

	total := math.Pow(float64(maxVal), float64(n))
	for i := n; i <= n*maxVal; i++ {
		fmt.Printf("%d, %f\n", i, float64(probabilities[flag][i])/total)
	}
}
