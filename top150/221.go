package top150

func maximalSquare(matrix [][]byte) int {

	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m)

	maxD := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxD = 1
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			if dp[i][j] > maxD {
				maxD = dp[i][j]
			}
		}
	}

	return maxD * maxD
}
