package top150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return recursion(root, p, q)
}

func recursion(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	l := recursion(root.Left, p, q)
	r := recursion(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	}
	return r
}
