package top150

func coinChange(coins []int, amount int) int {

	dp := []int{}
	for i := 0; i < amount+1; i++ {
		dp = append(dp, amount+1)
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if i >= v {
				dp[i] = min(dp[i], dp[i-v]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
