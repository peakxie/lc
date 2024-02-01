package top150

func kthSmallest(root *TreeNode, k int) int {

	res := []int{}

	var mid func(*TreeNode)
	mid = func(node *TreeNode) {
		if node == nil {
			return
		}

		if len(res) >= k {
			return
		}

		mid(node.Left)
		res = append(res, node.Val)
		mid(node.Right)
	}

	mid(root)
	if len(res) >= k {
		return res[k-1]
	}
	return -1
}
