package top150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return recursion(root, 0)
}

func recursion(root *TreeNode, pre int) int {
	if root == nil {
		return 0
	}

	sum := pre*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}

	return recursion(root.Left, sum) + recursion(root.Right, sum)
}
