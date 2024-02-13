package top150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	q := []*TreeNode{root}

	cc := 0
	for len(q) > 0 {
		cc++
		t := q
		q = nil
		tv := []int{}
		for len(t) > 0 {
			n := t[0]
			tv = append(tv, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			t = t[1:]
		}
		if cc%2 == 0 {
			reverse(tv)
		}
		res = append(res, tv)
	}
	return res
}

func reverse(list []int) {
	l, r := 0, len(list)-1
	for l < r {
		list[l], list[r] = list[r], list[l]
		l++
		r--
	}
}
