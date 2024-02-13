package top150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
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
		res = append(res, tv)
	}
	return res
}
