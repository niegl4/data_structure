package array

import "fmt"

/*
二十九
顺时针打印二位矩阵
*/
func printMatrixClockwise(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	columns := len(matrix[0])
	for _, row := range matrix {
		if len(row) != columns {
			return
		}
	}
	start := 0
	for rows > start * 2 && columns > start * 2 {
		printMatrixInCircle(matrix, start, rows, columns)
		start++
	}
	return
}

func printMatrixInCircle(matrix [][]int, start, rows, columns int) {
	jEnd := columns - 1 - start
	iEnd := rows - 1 - start
	var tmp []int
	for j := start; j <= jEnd; j++ {
		tmp = append(tmp, matrix[start][j])
	}
	if len(tmp) != 0 {
		fmt.Println(tmp)
	}
	tmp = tmp[0:0]

	if iEnd > start {
		for i := start+1; i <= iEnd; i++ {
			tmp = append(tmp, matrix[i][jEnd])
		}
	}
	if len(tmp) != 0 {
		fmt.Println(tmp)
	}
	tmp = tmp[0:0]

	if jEnd - start >= 1 && iEnd - start >= 1 {
		for j := jEnd-1; j >= start; j-- {
			tmp = append(tmp, matrix[iEnd][j])
		}
	}
	if len(tmp) != 0 {
		fmt.Println(tmp)
	}
	tmp = tmp[0:0]

	if iEnd - start >= 2 && jEnd - start >= 1 {
		for i := iEnd-1; i >= start+1; i-- {
			tmp = append(tmp, matrix[i][start])
		}
	}
	if len(tmp) != 0 {
		fmt.Println(tmp)
	}
	return
}
