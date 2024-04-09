package top150

func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])

	l, r := 0, m*n-1

	for l <= r {
		mid := (l + r) / 2
		i, j := pos(mid, n)
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}

func pos(p, n int) (int, int) {
	i := p / n
	j := p % n
	return i, j
}
