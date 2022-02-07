package array

/*
五十七-1
递增数组与一个和s。在数组中查找两个数，使得它们的和为s。任意一对满足要求即可。
*/
func findNumsWithSum(arr []int, sum int) (int, int, bool) {
	if len(arr) < 1 {
		return 0, 0, false
	}

	aheadIdx := 0
	behindIdx := len(arr) - 1
	for behindIdx > aheadIdx {
		s := arr[aheadIdx] + arr[behindIdx]
		if s == sum {
			return arr[aheadIdx], arr[behindIdx], true
		} else if s < sum {
			aheadIdx++
		} else {
			behindIdx--
		}
	}
	return 0, 0, false
}

/*
五十七-2
和为s的连续正数序列。
*/
func continuousSeq(sum int) (seqSet [][]int) {
	if sum < 3 {
		return nil
	}

	small := 1
	big := 2
	mid := (sum + 1) / 2
	curSum := small + big

	for small < mid {
		if curSum == sum {
			var tmp []int
			for i := small; i <= big; i++ {
				tmp = append(tmp, i)
			}
			seqSet = append(seqSet, tmp)
		}

		for curSum > sum && small < mid {
			curSum -= small
			small++

			if curSum == sum {
				var tmp []int
				for i := small; i <= big; i++ {
					tmp = append(tmp, i)
				}
				seqSet = append(seqSet, tmp)
			}
		}

		big++
		curSum += big
	}
	return seqSet
}
