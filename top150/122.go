package top150

func maxProfit(prices []int) int {
	n := len(prices)

	dp := make([][2]int, n)
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(i, j int) int { // W: func `max` is unused (unused)
	if i > j {
		return i
	}
	return j
}

// 第二种解法
func maxProfit(prices []int) int {
	n := len(prices)

	buy := -prices[0]
	sell := 0

	for i := 1; i < n; i++ {
		buy = max(buy, sell-prices[i])
		sell = max(sell, buy+prices[i])
	}
	return sell
}
