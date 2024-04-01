package top150

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {

	n := len(grid)

	var dfs func([][]int, int, int) *Node
	dfs = func(rows [][]int, c1, c2 int) *Node {
		for _, row := range rows {
			for i := c1; i < c2; i++ {
				if row[i] != rows[0][c1] {
					return &Node{
						Val:         false,
						IsLeaf:      false,
						TopLeft:     dfs(rows[:len(rows)/2], c1, (c1+c2)/2),
						TopRight:    dfs(rows[:len(rows)/2], (c1+c2)/2, c2),
						BottomLeft:  dfs(rows[len(rows)/2:], c1, (c1+c2)/2),
						BottomRight: dfs(rows[len(rows)/2:], (c1+c2)/2, c2),
					}
				}
			}
		}
		return &Node{Val: rows[0][c1] == 1, IsLeaf: true}
	}

	return dfs(grid, 0, n)
}
