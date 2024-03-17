package top150

func longestPalindrome(s string) string {

	n := len(s)

	if n < 2 {
		return s
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	maxLen := 1
	beginIndex := 0
	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := i + l - 1
			if j >= n {
				continue
			}
			if s[i] == s[j] {
				if j-i > 2 {
					dp[i][j] = dp[i+1][j-1]
				} else {
					dp[i][j] = 1
				}
			}

			if dp[i][j] > 0 && maxLen < j-i+1 {
				maxLen = j - i + 1
				beginIndex = i
			}
		}
	}
	return s[beginIndex : beginIndex+maxLen]
}
