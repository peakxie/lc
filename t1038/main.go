package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum(root, 0)
	return root
}

func sum(n *TreeNode, s int) int {
	if n == nil {
		return s
	}
	s = sum(n.Right, s)
	s += n.Val
	n.Val = s
	s = sum(n.Left, s)
	return s
}
