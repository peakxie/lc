package top150

import "math"

func maxSubArray(nums []int) int {

	maxV := math.MinInt / 2
	val := math.MinInt / 2

	for _, v := range nums {
		val = max(v, v+val)
		maxV = max(maxV, val)
	}

	return maxV
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
