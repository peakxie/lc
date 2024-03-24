package top150

func gameOfLife(board [][]int) {

	m, n := len(board), len(board[0])

	second := make([][]int, m)
	for i := 0; i < m; i++ {
		second[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			second[i][j] = board[i][j]
		}
	}

	position := []int{0, 1, -1}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			live := 0
			for ii := 0; ii < 3; ii++ {
				for jj := 0; jj < 3; jj++ {
					if !(ii == 0 && jj == 0) {
						r := i + position[ii]
						c := j + position[jj]
						if r >= 0 && r < m && c >= 0 && c < n && second[r][c] == 1 {
							live++
						}
					}
				}
			}

			if board[i][j] == 1 && (live < 2 || live > 3) {
				board[i][j] = 0
			}

			if board[i][j] == 0 && live == 3 {
				board[i][j] = 1
			}
		}
	}
}
