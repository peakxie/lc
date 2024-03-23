package top150

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	ml := make([]bool, m)
	nl := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				ml[i] = true
				nl[j] = true
			}
		}
	}

	for i, v := range matrix {
		for j, _ := range v {
			if ml[i] || nl[j] {
				v[j] = 0
			}
		}
	}
}
