package array

import "errors"

/*
四
二维递增数组，查找一个数字

O(rows+columns)

矩阵的右上角数字：一行中的最大，一列中的最小。
*/
func find(matrix [][]int, number int) (isFind bool, err error) {
	rows := len(matrix)
	if rows == 0 {
		return false, nil
	}
	columns := len(matrix[0])
	for i := 0; i < rows; i++ {
		if len(matrix[i]) != columns {
			return false, errors.New("matrix invalid")
		}
	}

	row := 0
	col := columns - 1
	for row < rows && col >= 0 {
		if matrix[row][col] == number {
			return true, nil
		}
		if matrix[row][col] < number {
			row++
		} else {
			col--
		}
	}
	return false, nil
}
