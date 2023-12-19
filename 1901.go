package main

import "math"

func findPeakGrid(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])

	maxind := func(s []int) int {
		ind := 0
		maxint := 0
		for i, v := range s {
			if v > maxint {
				maxint = v
				ind = i
			}
		}
		return ind
	}

	get := func(i, j int) int {
		if i == -1 || i == m {
			return math.MinInt
		}
		if j == -1 || j == n {
			return math.MinInt
		}
		return mat[i][j]
	}

	u, d := 0, m-1
	for {
		mid := (u + d) / 2
		j := maxind(mat[mid])
		if get(mid, j) > get(mid-1, j) && get(mid, j) > get(mid+1, j) {
			return []int{mid, j}
		}

		if get(mid+1, j) > get(mid, j) {
			u = mid + 1
		} else {
			d = mid - 1
		}
	}
}
