package top150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		res = append(res, q[len(q)-1].Val)
		t := q
		q = nil
		for len(t) > 0 {
			n := t[0]
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			t = t[1:]
		}
	}
	return res
}
