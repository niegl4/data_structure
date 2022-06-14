package array

/*
五-2
合并两个排序数组。要求空间复杂度为O(1)

从后向前遍历
*/
func mergeSort(arr1 *[]int, arr2 []int) {
	if arr1 == nil {
		return
	}

	len1 := len(*arr1)
	len2 := len(arr2)
	if len1 == 0 && len2 == 0 {
		return
	} else if len1 == 0 && len2 != 0 {
		*arr1 = append(*arr1, arr2...)
		return
	} else if len1 != 0 && len2 == 0 {
		return
	}

	for i := 0; i < len2; i++ {
		*arr1 = append(*arr1, 0)
	}

	index1 := len1 - 1
	index2 := len2 - 1
	newIndex := index1 + len2
	for index1 >= 0 && index2 >= 0 && newIndex >= 0 {
		if arr2[index2] > (*arr1)[index1] {
			(*arr1)[newIndex] = arr2[index2]
			index2--
		} else {
			(*arr1)[newIndex] = (*arr1)[index1]
			index1--
		}
		newIndex--
	}
	if index1 > 0 { //其实如果arr1有剩余，可以不用处理。因为是在arr1的基础上生成新arr。
		for i := index1; i >= 0; i-- {
			(*arr1)[newIndex] = (*arr1)[i]
			newIndex--
		}
	} else {
		for i := index2; i >= 0; i-- {
			(*arr1)[newIndex] = arr2[i]
			newIndex--
		}
	}
}
