package top150

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

	maxV := math.MinInt
	var getMaxV func(*TreeNode) int
	getMaxV = func(node *TreeNode) int {

		if node == nil {
			return 0
		}

		l := getMaxV(node.Left)
		l = max(l, 0)
		r := getMaxV(node.Right)
		r = max(r, 0)

		v := l + r + node.Val
		maxV = max(maxV, v)

		return node.Val + max(l, r)
	}

	getMaxV(root)
	return maxV
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
