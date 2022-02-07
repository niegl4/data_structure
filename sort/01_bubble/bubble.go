package _1_bubble

//严格来说，不是标准的冒泡排序算法。它不满足"两两比较相邻记录"的冒泡排序思想，仅仅就是最简单的交换排序而已。
//在每一次"冒泡"过程中，如果原本应该位于i处的元素，实际位于比较靠后的位置。这样就会造成原本应该位于i+1处的元素，经过本次遍历后，
//出现在那个靠后的位置。这会对下一次"冒泡"过程不利。
//也就是说当前次的"冒泡"，对下一次可能起不到什么作用。
//[X  ...    2    ...    1]  ==>  [1 ... ... 2]
func BubSort1(inputArr []int) {
	num := len(inputArr)
	if num <= 1 {
		return
	}

	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			// 这里要写大于号，而不是大于等于号。冒泡排序就是稳定的排序算法。
			if inputArr[i] > inputArr[j] {
				inputArr[i], inputArr[j] = inputArr[j], inputArr[i]
			}
		}
	}
}

//在每次"冒泡"过程中，都尽量让较小的元素向前移动。除了获取到i处元素外，还让整体更有序。
func BubSort2(inputArr []int) {
	num := len(inputArr)
	if num <= 1 {
		return
	}

	for i := 0; i < num; i++ {
		for j := num - 2; j >= i; j-- {
			if inputArr[j] > inputArr[j+1] {
				inputArr[j], inputArr[j+1] = inputArr[j+1], inputArr[j]
			}
		}
	}
}

//优化后的冒泡排序。如果某一次"冒泡"过程中，没有交换任何元素，那么说明这些元素已经是有序的，就可以及时结束排序。
//通过设置标志位，来实现判断。
func BubSort3(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	flag := true
	for i := 0; i < length && flag; i++ {
		flag = false
		for j := length - 2; j >= i; j-- { //总是从倒数第2个元素开始，从后往前。倒数第2，就是num-2.
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
	}
}
