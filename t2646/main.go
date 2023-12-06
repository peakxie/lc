package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	ll := make([][]int, n)
	for _, v := range edges {
		ll[v[0]] = append(ll[v[0]], v[1])
		ll[v[1]] = append(ll[v[1]], v[0])
	}

	count := make([]int, n)
	var dfs func(int, int, int) bool
	dfs = func(node, p, end int) bool {
		if node == end {
			count[node]++
			return true
		}
		for _, v := range ll[node] {
			if v == p {
				continue
			}
			if dfs(v, node, end) {
				count[node]++
				return true
			}
		}
		return false
	}

	for _, v := range trips {
		dfs(v[0], -1, v[1])
	}

	var dp func(int, int) []int
	dp = func(node, p int) []int {
		res := []int{
			price[node] * count[node], price[node] * count[node] / 2,
		}
		for _, v := range ll[node] {
			if v == p {
				continue
			}
			vv := dp(v, p)
			res[0] = res[0] + min(vv[0], vv[1])
			res[1] = res[1] + vv[0]
		}
		return res
	}
	// 不知道为啥以0为跟节点计算
	res := dp(0, -1)
	return min(res[0], res[1])
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
