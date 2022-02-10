package search

import "testing"

func TestSearchMinInRotateSortArr1(t *testing.T) {
	testSet := [][]int{
		{1, 1, 1, 1, 1, 1},
	}
	resultSet := []int{
		0,
	}
	for i, target := range testSet {
		searchRes := searchMinInRotateSortArr1(target)
		if resultSet[i] != searchRes {
			t.Fatal("failed", target)
		} else {
			t.Log(resultSet[i])
		}
	}
}

func TestSearchMinInRotateSortArr2(t *testing.T) {
	testSet := [][]int{
		{3, 4, 5, 1, 2},
		{1, 2, 3, 4, 5},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1},
	}
	resultSet := []int{
		3,
		0,
		1,
		3,
	}
	for i, target := range testSet {
		searchRes := searchMinInRotateSortArr2(target)
		if resultSet[i] != searchRes {
			t.Fatal("failed", target)
		} else {
			t.Log(resultSet[i])
		}
	}
}
