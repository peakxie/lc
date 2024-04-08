package top150

import "math"

func maxSubarraySumCircular(nums []int) int {

	minV := math.MaxInt / 2
	minv := math.MaxInt / 2
	maxV := math.MinInt / 2
	maxv := math.MinInt / 2
	total := 0

	for _, v := range nums {
		total += v

		maxv = max(v, v+maxv)
		maxV = max(maxV, maxv)

		minv = min(v, v+minv)
		minV = min(minV, minv)
	}

	if total == minV {
		return maxV
	}

	return max(maxV, total-minV)
}

func max(i, j int) int { // E: max redeclared in this block
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int { // E: max redeclared in this block
	if i < j {
		return i
	}
	return j
}
