package sort

var (
	Sli1 = []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	Sli2 = []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	//适用于改进版bubble的数据
	Sli3 = []int{2, 1, 3, 4, 5, 6, 7, 8, 9}
)

func CheckSlice(sli []int) bool {
	n := len(sli)
	if n < 2 {
		return true
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if sli[i] > sli[j] {
				return false
			}
		}
	}
	return true
}
