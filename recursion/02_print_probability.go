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

func printProbabilityV1(n int) {
	if n < 1 {
		return
	}

	maxSum := n * maxVal
	probabilities := make([]int, maxSum-n+1)

	probability(n, probabilities)

	total := math.Pow(float64(n), float64(maxVal))
	for i := n; i <= maxSum; i++ {
		fmt.Printf("%d, %f\n", i, float64(probabilities[i-n])/total)
	}
}

func probability(n int, probabilities []int) {
	for i := 1; i <= maxVal; i++ {
		probabilityCore(n, n, i, probabilities)
	}
}

func probabilityCore(original, current, sum int, probabilities []int) {
	if current == 1 {
		probabilities[sum-original]++
	} else {
		for i := 1; i <= maxVal; i++ {
			probabilityCore(original, current-1, sum+i, probabilities)
		}
	}
}

func printProbabilityV2(n int) {
	if n < 1 {
		return
	}

	probabilities := make([][]int, 2)
	probabilities[0] = make([]int, maxVal*n+1)
	probabilities[1] = make([]int, maxVal*n+1)
	flag := 0

	for i := 1; i <= maxVal; i++ {
		probabilities[flag][i] = 1
	}

	for k := 2; k <= n; k++ {
		for i := 0; i < k; i++ {
			probabilities[1-flag][i] = 0
		}

		for i := k; i <= maxVal*k; i++ {
			probabilities[1-flag][i] = 0
			for j := 1; j <= i && j <= maxVal; j++ {
				probabilities[1-flag][i] += probabilities[flag][i-j]
			}
		}

		flag = 1 - flag
	}

	total := math.Pow(float64(n), float64(maxVal))
	for i := n; i <= n*maxVal; i++ {
		fmt.Printf("%d, %f\n", i, float64(probabilities[flag][i])/total)
	}
}
