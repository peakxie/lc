package top150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	list := recursion(root)
	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left, pre.Right = nil, cur
	}
}

func recursion(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	list := []*TreeNode{}
	list = append(list, root)
	list = append(list, recursion(root.Left)...)
	list = append(list, recursion(root.Right)...)

	return list
}
