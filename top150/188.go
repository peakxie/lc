package top150

// 有一个用例过不了,再看看
// prices =
// [5,7,2,7,3,3,5,3,0]
// 输出
// 10
// 预期结果
// 9
func maxProfit(k int, prices []int) int {

	n := len(prices)

	buy := make([]int, k)
	sell := make([]int, k)

	for i := 0; i < k; i++ {
		buy[i] = -prices[i%n]
	}

	for i := 1; i < n; i++ {
		buy[0] = max(buy[0], -prices[i])
		sell[0] = max(sell[0], buy[0]+prices[i])
		for j := 1; j < k; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}

	return sell[k-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
