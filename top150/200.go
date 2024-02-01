package top150

func numIslands(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])

	res := 0
	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	var dfs func(int, int)
	dfs = func(i, j int) {
		if !valid(m, n, i, j) {
			return
		}
		if grid[i][j] != '1' {
			return
		}

		grid[i][j] = '2'

		for _, v := range dirs {
			dfs(i+v[0], j+v[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)

			}
		}
	}

	return res

}

func valid(m, n, i, j int) bool {
	if i >= 0 && i < m && j >= 0 && j < n {
		return true
	}
	return false
}
