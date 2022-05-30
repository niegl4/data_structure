package array

import (
	"errors"
)

/*
四十七
m*n的二维矩阵，每个位置放有一个礼物。从左上角开始，每次向下或向右移动一步，直到右下角。礼物的最大价值是多少。
*
递归分析问题，循环实现代码。动态规划问题。
*/

func maxValue(matrix [][]int) (maxVal int, err error) {
	rows := len(matrix)
	if rows < 1 {
		return 0, nil
	}
	columns := len(matrix[0])
	for _, col := range matrix {
		if columns != len(col) {
			return 0, errors.New("matrix invalid")
		}
	}

	values := make([]int, columns)
	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			left := 0
			up := 0

			if row > 0 {
				up = values[col]
			}

			if col > 0 {
				left = values[col-1]
			}

			if left > up {
				values[col] = left
			} else {
				values[col] = up
			}
			values[col] += matrix[row][col]

		}
	}

	return values[columns-1], nil
}
