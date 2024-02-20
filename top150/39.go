package top150

func combinationSum(candidates []int, target int) [][]int {

	ans := [][]int{}

	var dfs func(int, []int)
	dfs = func(i int, item []int) {
		sum := 0
		for _, v := range item {
			sum += v
		}
		if sum == target {
			dst := make([]int, len(item))
			copy(dst, item)
			ans = append(ans, dst)
			return
		}

		if sum > target {
			return
		}

		for j := i; j < len(candidates); j++ {
			item = append(item, candidates[j])
			dfs(j, item)
			item = item[:len(item)-1]
		}

	}

	dfs(0, []int{})
	return ans
}
