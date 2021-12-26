package _1_bi_tree

import "testing"

func TestMarshalBT(t *testing.T) {
	root := &BiTree{
		Data: 1,
		lChild: &BiTree{
			Data: 2,
			lChild: &BiTree{
				Data:   4,
				lChild: nil,
				rChild: nil,
			},
			rChild: nil,
		},
		rChild: &BiTree{
			Data: 3,
			lChild: &BiTree{
				Data:   5,
				lChild: nil,
				rChild: nil,
			},
			rChild: &BiTree{
				Data:   6,
				lChild: nil,
				rChild: nil,
			},
		},
	}
	res := marshalBT(root)
	t.Log(res)

	root = unmarshalBT([]interface{}{1, 2, 4, "nil", "nil", "nil", 3, 5, "nil", "nil", 6, "nil", "nil"})
	res = marshalBT(root)
	t.Log(res)
}
