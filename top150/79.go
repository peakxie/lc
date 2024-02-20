package top150

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var dfs func(int, int, int) bool
	dfs = func(x, y, i int) bool {
		if i == len(word) {
			return true
		}
		if x < 0 || x >= m || y < 0 || y >= n {
			return false
		}
		ch := word[i]
		if board[x][y] != ch {
			return false
		}

		board[x][y] = '*'
		for _, v := range dirs {
			if dfs(x+v[0], y+v[1], i+1) {
				return true
			}
		}
		board[x][y] = ch

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
