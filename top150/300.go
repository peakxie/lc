package top150

func lengthOfLIS(nums []int) int {

	n := len(nums)
	dp := []int{}

	for i := 0; i < n; i++ {
		dp = append(dp, 1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 1
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
