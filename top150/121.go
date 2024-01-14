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

// 第二种解法
func maxProfit(prices []int) int {
	n := len(prices)

	buy := -prices[0]
	sell := 0

	for i := 1; i < n; i++ {
		buy = max(buy, -prices[i])
		sell = max(sell, buy+prices[i])
	}

	return sell
}
