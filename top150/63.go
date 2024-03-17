package top150

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([]int, n)

	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[n-1]
}
