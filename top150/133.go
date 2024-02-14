package top150

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {

	m := map[int]*Node{}

	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		if n, ok := m[node.Val]; ok {
			return n
		}

		newN := &Node{
			Val: node.Val,
		}
		m[newN.Val] = newN

		for _, v := range node.Neighbors {
			n := dfs(v)
			newN.Neighbors = append(newN.Neighbors, n)
		}

		return newN
	}

	return dfs(node)
}
