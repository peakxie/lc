package top150

import "math"

func removeDuplicates(nums []int) int {
	n := len(nums)
	l := 0
	cur := math.MinInt
	for i := 0; i < n; i++ {
		if cur != nums[i] {
			cur = nums[i]
			nums[l] = nums[i]
			l++
		}
	}
	return l
}
