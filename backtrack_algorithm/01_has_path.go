package backtrack_algorithm

import (
	"errors"
)

/*
十二
*
判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。

回溯问题，递归代码处理，通常分为主函数，递归核心函数。
这里，主函数是逐个判断矩阵中的点。核心函数是用来从指定点出发探索路径。
*/

func hasPath(matrix [][]byte, str string) (bool, error) {
	rows := len(matrix)
	if rows == 0 {
		return false, errors.New("matrix invalid")
	}
	columns := len(matrix[0])
	for _, col := range matrix {
		if len(col) != columns {
			return false, errors.New("matrix invalid")
		}
	}
	if len(str) == 0 {
		return false, errors.New("string invalid")
	}

	matrixVisited := make([][]bool, rows)
	for idx := range matrixVisited {
		matrixVisited[idx] = make([]bool, columns)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			if hasPathCore(matrix, row, col, str, 0, matrixVisited) {
				return true, nil
			}
		}
	}
	return false, nil
}

func hasPathCore(matrix [][]byte, row, col int, str string, strIndex int, matrixVisited [][]bool) bool {
	if strIndex >= len(str) {
		return true
	}

	hasPath := false
	if row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0]) &&
		matrix[row][col] == str[strIndex] && !matrixVisited[row][col] {
		strIndex++
		matrixVisited[row][col] = true
		hasPath = hasPathCore(matrix, row-1, col, str, strIndex, matrixVisited) ||
			hasPathCore(matrix, row, col-1, str, strIndex, matrixVisited) ||
			hasPathCore(matrix, row+1, col, str, strIndex, matrixVisited) ||
			hasPathCore(matrix, row, col+1, str, strIndex, matrixVisited)

		if !hasPath {
			matrixVisited[row][col] = false //visited是一个引用类型的矩阵，本次判断为noPath后，要【恢复visited】
			//strIndex--//strIndex是一个值类型的索引，每次调用core函数都会传递新值，不用恢复
		}
	}
	return hasPath
}
