package _2_ring

import (
	ring2 "container/ring"
)

/*
六十二
0~n-1，n个数排成一个圆圈。从数字0开始，每次删除第m个数字，求圆圈剩余的最后一个数字。
*/
func lastRemainV1(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	}

	ring := ring2.New(n)
	for i := 0; i < n; i++ {
		ring.Value = i
		ring = ring.Next()
	}

	for ring.Len() > 1 {
		for i := 0; i < m-1; i++ {
			ring = ring.Move(1)
		}
		ring = ring.Prev()
		//fmt.Println(ring.Next())
		ring.Unlink(1)
		ring = ring.Next()
	}
	return ring.Value.(int)
}

func lastRemainV2(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	}

	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
