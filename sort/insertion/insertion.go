package insertion

//不断将一个记录插入到前面已经排好序的有序列表中
func InsertionSort(input []int) {
	n := len(input)
	if n <= 1 {
		return
	}

	var j int
	//n长的数组，只用循环n-1次
	for i := 0; i < n-1; i++ {

		//当前>后一个，再进入循环
		if input[i] > input[i+1] {
			//记录这个后一个的值
			temp := input[i+1]
			//从当前直到第0个，逐个与temp比，比temp大就向后填充。todo：没必要从j=i开始吧，从i-1开始
			for j = i; j >= 0 && input[j] > temp; j-- {
				input[j+1] = input[j]
			}
			//j的作用域要大于内层循环，所以提前声明
			input[j+1] = temp
		}
	}
}
