package top150

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := [2]int{0, nums[0]}

	for i := 1; i < n; i++ {
		t := dp[0]
		dp[0] = max(dp[0], dp[1])
		dp[1] = t + nums[i]
	}

	return max(dp[0], dp[1])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
