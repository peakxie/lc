package top150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	vals []int
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{}
	bst.inorder(root)
	return bst
}

func (this *BSTIterator) inorder(root *TreeNode) {
	if root == nil {
		return
	}
	this.inorder(root.Left)
	this.vals = append(this.vals, root.Val)
	this.inorder(root.Right)
}

func (this *BSTIterator) Next() int {
	val := this.vals[0]
	this.vals = this.vals[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.vals) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
