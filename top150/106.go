package top150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	rv := postorder[len(postorder)-1]
	root := &TreeNode{
		Val: rv,
	}

	n := len(inorder)
	i := 0
	for ; i < n; i++ {
		if inorder[i] == rv {
			break
		}
	}

	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])

	return root

}
