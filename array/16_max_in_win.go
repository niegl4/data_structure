package array

/*
五十九-1
数组中，滑动窗口中的最大值。
****
*/
func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	//从队尾逐个比较，直到队列空或者一个元素大于新元素
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	//起始部分，滑动窗口未满
	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)           //push本身就是在维护递减队列
		for q[0] <= i-k { //窗口移动，可能把队列头部的元素移出
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
