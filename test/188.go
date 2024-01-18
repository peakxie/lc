package main

import "fmt"

func main() {
	prices := []int{5, 7, 2, 7, 3, 3, 5, 3, 0}
	k := 4
	res := maxProfit(k, prices)
	fmt.Println(res)
}

// 有一个用例过不了,再看看
// prices =
// [5,7,2,7,3,3,5,3,0]
// k = 4
// 输出
// 10
// 预期结果
// 9
func maxProfit(k int, prices []int) int {

	n := len(prices)

	k = min(n, k)

	buy := make([]int, k)
	sell := make([]int, k)

	for i := 0; i < k; i++ {
		buy[i] = -prices[0]
	}

	fmt.Printf("第%d天:\n", 0)
	for ii := 0; ii < k; ii++ {
		fmt.Printf("buy[%d](%d), sell[%d](%d)\n", ii, buy[ii], ii, sell[ii])
	}

	for i := 1; i < n; i++ {
		buy[0] = max(buy[0], -prices[i])
		sell[0] = max(sell[0], buy[0]+prices[i])
		for j := 1; j < k; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}

		fmt.Printf("第%d天:\n", i)
		for ii := 0; ii < k; ii++ {
			fmt.Printf("buy[%d](%d), sell[%d](%d)\n", ii, buy[ii], ii, sell[ii])
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

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
