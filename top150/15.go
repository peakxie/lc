package top150

import "sort"

func threeSum(nums []int) [][]int {

	n := len(nums)
	res := [][]int{}

	if n < 3 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < n; i++ {

		if nums[i] > 0 {
			return res
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		// TODO 调试多次后才成功
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
