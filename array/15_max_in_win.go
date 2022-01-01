package array

import "errors"

/*
五十九-1
数组中，滑动窗口中的最大值。
****
*/
func maxInWindows(arr []int, size int) (int, error) {
	length := len(arr)
	if length == 0 {
		return 0, errors.New("input invalid")
	}

	win := make([]int, 0, size)
	maxInWin := 0
	for i := 0; i < length; i++ {
		if len(win) == 0 {
			win = append(win, i)
			if maxInWin < arr[i] {
				maxInWin = arr[i]
			}
			continue
		}

		//滑动窗口移动，过期下标划出
		if len(win) == size {
			minIdxInWin := i - size + 1
			// win[0] < 允许出现的最小下标 <= win[len-1]
			if minIdxInWin > win[0] && minIdxInWin <= win[len(win)-1] {
				j := 0
				for ; j < len(win); j++ {
					if win[j] == minIdxInWin {
						break
					}
				}
				win = win[j:]
			} else if minIdxInWin > win[len(win)-1] { // 允许出现的最小下标 > win[len-1]
				win = win[0:0]
				continue
			}
		}

		if arr[i] >= arr[win[0]] { // >=窗口中的最大值，直接更新
			win = []int{i}
		} else {
			if arr[i] <= arr[win[len(win)-1]] { // <=窗口中的最小值，直接append
				win = append(win, i)
			} else {
				for idx := range win { // 最小值 < < 最大值，截取
					if arr[win[idx]] < arr[i] {
						win = win[0:idx]
						win = append(win, i)
					}
				}
			}
		}

		if len(win) > 0 && maxInWin < arr[win[0]] {
			maxInWin = arr[win[0]]
		}
	}
	return maxInWin, nil
}
