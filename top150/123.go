package top150

func maxProfit(prices []int) int {

	n := len(prices)

	buy := [2]int{-prices[0], -prices[0]}
	sell := [2]int{}

	for i := 0; i < n; i++ {
		buy[0] = max(buy[0], -prices[i])
		sell[0] = max(sell[0], buy[0]+prices[i])
		for j := 1; j < len(buy); j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}

	return sell[1]

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
