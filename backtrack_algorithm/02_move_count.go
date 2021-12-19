package backtrack_algorithm

import "errors"

/*
十三
m*n的矩阵，从（0，0）开始移动，不能进入横纵坐标的数位之和大于k的坐标。能到达多少个坐标点？

 */
func moveCount(threshold, rows, columns int) (int, error) {
	if threshold < 0 || rows <= 0 || columns <= 0 {
		return 0, errors.New("input invalid")
	}

	matrixVisited := make([][]bool, rows)
	for idx := range matrixVisited {
		matrixVisited[idx] = make([]bool, columns)
	}
	return moveCountCore(threshold, rows, columns, 0, 0, matrixVisited), nil

}

func moveCountCore(threshold, rows, columns, row, col int, matrixVisited [][]bool) int {
	count := 0
	if check(threshold, rows, columns, row, col, matrixVisited) {
		matrixVisited[row][col] = true
		count = 1 + moveCountCore(threshold, rows, columns, row-1, col, matrixVisited) +
			moveCountCore(threshold, rows, columns, row, col-1, matrixVisited) +
			moveCountCore(threshold, rows, columns, row+1, col, matrixVisited) +
			moveCountCore(threshold, rows, columns, row, col+1, matrixVisited)
	}
	return count
}

func check(threshold, rows, columns, row, col int, matrixVisited [][]bool) bool {
	return row >= 0 && row < rows && col >= 0 && col < columns &&
		!matrixVisited[row][col] && getDigitSum(row) + getDigitSum(col) <= threshold
}

func getDigitSum(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}
	return sum
}