package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	minind := 0xffffff
	for i := 0; i < m; i++ {
		ind := 0
		for j := 0; j < n; j++ {
			for k := 0; k < m*n; k++ {
				if arr[k] == mat[i][j] {
					ind = max(ind, k)
					break
				}
			}
		}
		minind = min(minind, ind)
	}

	for j := 0; j < n; j++ {
		ind := 0
		for i := 0; i < m; i++ {
			for k := 0; k < m*n; k++ {
				if arr[k] == mat[i][j] {
					ind = max(ind, k)
					break
				}
			}
		}
		minind = min(minind, ind)
	}

	return minind

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
