package array

import (
	"fmt"
	"testing"
)

func TestPrintMatrixClockwise(t *testing.T) {
	matrixSet := [][][]int{
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
		{
			{1},
			{5},
			{7},
		},
		{
			{1, 2, 3, 4},
		},
	}
	for _, matrix := range matrixSet {
		printMatrixClockwise(matrix)
		fmt.Println("---------")
	}
}
