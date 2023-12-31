package top150

func maxProfit(prices []int) int {
	n := len(prices)
	minV := prices[0]
	maxV := 0
	for i := 1; i < n; i++ {
		maxV = max(maxV, prices[i]-minV)
		minV = min(minV, prices[i])
	}

	return maxV
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
