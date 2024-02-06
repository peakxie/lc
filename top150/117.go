package top150

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {

	queue := []*Node{root}
	for len(queue) > 0 {
		q := queue
		queue = nil
		for i := 0; i < len(q); i++ {
			if q[i] == nil {
				continue
			}
			if i+1 < len(q) {
				q[i].Next = q[i+1]
			}
			if q[i].Left != nil {
				queue = append(queue, q[i].Left)
			}
			if q[i].Right != nil {
				queue = append(queue, q[i].Right)
			}
		}
	}
	return root
}
