package top150

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)

	if l1*l2 == 0 {
		return l1 + l2
	}

	dp := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 0; i < l1+1; i++ {
		dp[i][0] = i
	}

	for i := 0; i < l2+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			a := dp[i-1][j] + 1
			b := dp[i][j-1] + 1
			c := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				c += 1
			}

			dp[i][j] = min(min(a, b), c)
		}
	}

	return dp[l1][l2]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
