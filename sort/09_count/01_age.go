package _9_count

import "errors"

/*
十一-0
对员工年龄排序，几万员工，要求时间复杂度O(n)，空间复杂度O(n)。
数值的大小范围是在一个比较小的范围。
 */
func sortAge(ages []int) error {
	if len(ages) <= 1 {
		return nil
	}

	maxAge := 99
	ageCount := make([]int, maxAge+1)
	for _, age := range ages {
		if age > maxAge || age <= 0 {
			return errors.New("age invalid")
		}
		ageCount[age]++
	}

	index := 0
	for age, count := range ageCount {
		for j := 0; j < count; j++ {
			ages[index] = age
			index++
		}
	}
	return nil
}