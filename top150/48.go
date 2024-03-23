package top150

func rotate(matrix [][]int) {

	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			// matrix[j][n-i-1] = matrix[i][j]
			matrix[j][n-i-1] = tmp
		}
	}
}
