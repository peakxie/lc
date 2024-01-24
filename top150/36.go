package top150

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]int{}
	cols := [9][9]int{}
	ss := [3][3][9]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ch := board[i][j]
			if ch != '.' {
				v := ch - '1'
				rows[i][v]++
				cols[j][v]++
				ss[i/3][j/3][v]++

				if rows[i][v] > 1 || cols[j][v] > 1 || ss[i/3][j/3][v] > 1 {
					return false
				}
			}
		}
	}
	return true
}
