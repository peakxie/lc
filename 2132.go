package main

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		u, d, l, r := 0, 0, 100000, 100000
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				l = min(l, j)
				r = min(r, n-1-j)
				if l > stampWidth-1 && i > stampHeight-1 {
					return true
				}
				if r > stampWidth-1 && i > stampHeight-1 {
					return true
				}
			}
		}
	}

	return false
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
