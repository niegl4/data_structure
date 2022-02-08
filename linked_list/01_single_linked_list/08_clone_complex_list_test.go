package _1_single_linked_list

import "testing"

/*
     ------------->
    |             |
 ------->        |
|   |   |       |
A - B - C - D - E
    |       |
	<-------
 */
func TestCloneComplexList(t *testing.T) {
	head := &ComplexList{
		Val: "A",
		Next: &ComplexList{
			Val: "B",
			Next: &ComplexList{
				Val: "C",
				Next: &ComplexList{
					Val: "D",
					Next: &ComplexList{
						Val:     "E",
						Next:    nil,
						Sibling: nil,
					},
					Sibling: nil,
				},
				Sibling: nil,
			},
			Sibling: nil,
		},
		Sibling: nil,
	}
	head.Sibling = head.Next.Next
	head.Next.Sibling = head.Next.Next.Next.Next
	head.Next.Next.Next.Sibling = head.Next

	newHead := cloneComplexList(head)
	node := newHead
	tmp := make([]interface{}, 0)
	for node != nil {
		tmp = append(tmp, node.Val)
		node = node.Next
	}
	t.Log(tmp)

	node = newHead
	tmp = tmp[0:0]
	for node != nil {
		if node.Sibling == nil {
			tmp = append(tmp, "nil")
		} else {
			tmp = append(tmp, node.Sibling.Val)
		}
		node = node.Next
	}
	t.Log(tmp)
}
