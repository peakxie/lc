package top150

func wordBreak(s string, wordDict []string) bool {

	n := len(s)
	dp := make([]bool, n+1)
	m := map[string]bool{}

	for _, v := range wordDict {
		m[v] = true
	}

	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
