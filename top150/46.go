package top150

func permute(nums []int) [][]int {

	res := [][]int{}

	var dfs func(int, []int)
	dfs = func(ind int, item []int) {
		if ind == len(nums) {
			dst := make([]int, len(nums))
			copy(item, dst)
			res = append(res, dst)
			return
		}

		for i := ind; i < len(nums); i++ {
			nums[ind], nums[i] = nums[i], nums[ind]
			dfs(ind+1, append(item, nums[ind]))
			nums[i], nums[ind] = nums[ind], nums[i]
		}

	}
	return res
}
