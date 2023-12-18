package main

import "math"

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	get := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt
		}
		return nums[i]

	}

	for {
		mid := (l + r) / 2
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return mid
		}

		if get(mid) < get(mid+1) {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
}
