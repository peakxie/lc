package top150

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])

	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if !valid(i, j, m, n) {
			return
		}

		if board[i][j] != 'O' {
			return
		}

		board[i][j] = 'A'
		for _, v := range dirs {
			dfs(i+v[0], j+v[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'A':
				board[i][j] = 'O'
			}
		}
	}
}

func valid(i, j, m, n int) bool {
	if i >= 0 && i < m && j >= 0 && j < n {
		return true
	}
	return false
}
