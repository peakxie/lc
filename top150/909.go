package top150

func trans(id, n int) (int, int) {
	r, c := (id-1)/n, (id-1)%n
	if r%2 == 1 {
		c = n - 1 - c
	}
	r = n - 1 - r
	return r, c
}

func snakesAndLadders(board [][]int) int {
	n := len(board)

	vis := make([]bool, n*n+1)

	queue := [][]int{{1, 0}}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for i := 1; i <= 6; i++ {
			next := v[0] + i
			if next > n*n {
				break
			}
			r, c := trans(next, n)
			if board[r][c] != -1 {
				next = board[r][c]
			}

			if next == n*n {
				return v[1] + 1
			}
			if !vis[next] {
				vis[next] = true
				queue = append(queue, []int{next, v[1] + 1})
			}
		}
	}

	return -1
}
