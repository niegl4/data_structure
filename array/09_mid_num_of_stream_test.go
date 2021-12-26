package array

import "testing"

func TestCalMidNum(t *testing.T) {
	numberStream := make(chan int, 0)
	go func() {
		arr := []int{1, 3, 3, 4, 5, 6, 7}
		for _, num := range arr {
			numberStream <- num
		}
	}()

	calMidNum(numberStream)
}
