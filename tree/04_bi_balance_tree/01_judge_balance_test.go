package _4_bi_balance_tree

import "testing"

func TestJudgeBalance(t *testing.T) {
	root := &BiTree{
		Data: 5,
		//LChild: &BiTree{
		//	Data: 3,
		//	LChild: &BiTree{
		//		Data:   2,
		//		LChild: nil,
		//		RChild: nil,
		//	},
		//	RChild: &BiTree{
		//		Data:   4,
		//		LChild: nil,
		//		RChild: nil,
		//	},
		//},
		RChild: &BiTree{
			Data: 7,
			LChild: &BiTree{
				Data:   6,
				LChild: nil,
				RChild: nil,
			},
			RChild: &BiTree{
				Data:   8,
				LChild: nil,
				RChild: nil,
			},
		},
	}
	t.Log(judgeBalance(root))
}
