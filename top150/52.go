package top150

func totalNQueens(n int) int {
	cols := make([]bool, n)
	x1s := make([]bool, 2*n)
	x2s := make([]bool, 2*n)

	res := 0
	var do func(row int)
	do = func(row int) {
		if row == n {
			res++
			return
		}

		for i := 0; i < n; i++ {
			x1 := row + n - 1 - i
			x2 := row + i
			if cols[i] || x1s[x1] || x2s[x2] {
				continue
			}
			cols[i] = true
			x1s[x1] = true
			x2s[x2] = true
			do(row + 1)
			cols[i] = false
			x1s[x1] = false
			x2s[x2] = false
		}
	}
	do(0)
	return res
}
