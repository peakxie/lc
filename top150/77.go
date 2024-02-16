package top150

func combine(n int, k int) (ans [][]int) {

	var dfs func(int, []int)
	dfs = func(i int, item []int) {
		if len(item) == k {
			dst := make([]int, k)
			copy(dst, item)
			ans = append(ans, dst)
			return
		}

		if len(item)+n-i+1 < k {
			return
		}

		for j := i; j <= n && i+1 != j; j++ {
			dfs(i+1, item)
			dfs(i+1, append(item, j))
		}

	}

	dfs(1, []int{})
	return ans

}
