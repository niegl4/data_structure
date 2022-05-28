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
	for rows > start*2 && columns > start*2 {
		printMatrixInCircle(matrix, start, rows, columns)
		start++
	}
	return
}

func printMatrixInCircle(matrix [][]int, start, rows, columns int) {
	jEnd := columns - 1 - start //行结束索引
	iEnd := rows - 1 - start    //列结束索引
	var tmp []int

	//总有第一行
	for j := start; j <= jEnd; j++ {
		tmp = append(tmp, matrix[start][j])
	}
	if len(tmp) != 0 {
		fmt.Println(tmp)
	}
	tmp = tmp[0:0]

	//列结束索引大于起始索引，就有第一列
	if iEnd > start {
		for i := start + 1; i <= iEnd; i++ {
			tmp = append(tmp, matrix[i][jEnd])
		}
	}
	if len(tmp) != 0 {
		fmt.Println(tmp)
	}
	tmp = tmp[0:0]

	//行结束索引大于起始 && 列结束索引大于起始
	if jEnd > start && iEnd > start {
		for j := jEnd - 1; j >= start; j-- {
			tmp = append(tmp, matrix[iEnd][j])
		}
	}
	if len(tmp) != 0 {
		fmt.Println(tmp)
	}
	tmp = tmp[0:0]

	//列结束索引比起始大于等于2 && 行结束索引大于起始
	if iEnd-start >= 2 && jEnd > start {
		for i := iEnd - 1; i >= start+1; i-- {
			tmp = append(tmp, matrix[i][start])
		}
	}
	if len(tmp) != 0 {
		fmt.Println(tmp)
	}
	return
}
