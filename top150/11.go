package top150

import "math"

func maxArea(height []int) int {
	l, r := 0, len(height)-1

	res := math.MinInt
	for l < r {
		v := -max(-height[l], -height[r]) * (r - l)
		res = max(res, v)
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
