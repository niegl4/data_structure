package _2_insertion

//不断将一个记录插入到前面已经排好序的有序列表中
func InsertSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i <= length-2; i++ { //要拿i与i+1比较，所以i只能到倒数第2。倒数第2，就是n-2.
		//当前>后一个，再进入循环
		if arr[i] > arr[i+1] {
			//记录这个最新出现的小元素
			temp := arr[i+1]
			//j的作用域要大于内层循环，所以提前声明。
			j := i
			//从当前直到第0个，逐个与temp比，比temp大就向后填充。
			for ; j >= 0 && arr[j] > temp; j-- {
				arr[j+1] = arr[j]
			}
			//循环退出的时候，无论是j<0，还是最新出现的那个小元素比有序序列大，j+1就是小元素的位置。
			arr[j+1] = temp
		}
	}
}
