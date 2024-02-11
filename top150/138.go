package top150

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	m := map[*Node]*Node{}
	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		if v, ok := m[n]; ok {
			return v
		}

		newN := &Node{
			Val: n.Val,
		}
		m[n] = newN
		newN.Next = dfs(n.Next)
		newN.Random = dfs(n.Random)
		return newN
	}
	nHead := dfs(head)
	return nHead
}
